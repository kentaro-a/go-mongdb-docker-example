package sessions

import (
	"app/models"
	"context"
	"crypto/rand"
	"errors"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	CookieName                   string        = "go-tutorial"
	SessionTimeout               time.Duration = 60 * 10
	DeleteExpiredSessionInterval time.Duration = 60 * 60
)

func init() {
	go clearExpiredSessions()
}

func clearExpiredSessions() {
	for {
		conn := models.GetConnection()
		defer func() {
			if err := conn.Disconnect(context.TODO()); err != nil {
				log.Fatal(err)
			}
		}()
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		collection := conn.Database("test").Collection("sessions")
		filter := bson.M{
			"expire_dt": bson.M{
				"$lt": time.Now(),
			},
		}

		_, err := collection.DeleteMany(ctx, filter)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * DeleteExpiredSessionInterval)
	}
}

func newSession(data map[string]interface{}) *Session {
	insert_dt := time.Now()
	expire_dt := insert_dt.Add(time.Second * SessionTimeout)
	sess_id, _ := MakeSessionID(64)
	sess := &Session{
		sess_id,
		data,
		insert_dt,
		expire_dt,
	}

	conn := models.GetConnection()
	defer func() {
		if err := conn.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := conn.Database("test").Collection("sessions")
	_, err := collection.InsertOne(ctx, sess)
	if err != nil {
		log.Fatal(err)
	}
	return sess
}

type Session struct {
	ID       string                 `json:"id" bson:"_id"`
	Data     map[string]interface{} `json:"data" bson:"data"`
	InsertDt time.Time              `json:"insert_dt" bson:"insert_dt"`
	ExpireDt time.Time              `json:"expire_dt" bson:"expire_dt"`
}

func (sess *Session) Write(key string, v interface{}) {
	sess.Data[key] = v
}
func (sess *Session) Read(key string) (interface{}, bool) {
	if v, ok := sess.Data[key]; ok {
		return v, true
	} else {
		return nil, false
	}
}

func IsExpiredSession(sess *Session) bool {
	is_expired := false
	if sess.ExpireDt.Before(time.Now()) {
		is_expired = true
	}
	return is_expired
}

func CheckSession(r *http.Request) (bool, *Session) {
	if sess, err := GetSession(r); err == nil {
		if IsExpiredSession(sess) {
			return false, nil
		} else {
			return true, sess
		}
	} else {
		return false, nil
	}
}

func GetSession(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return nil, err
	}
	sess_id := cookie.Value

	conn := models.GetConnection()
	defer func() {
		if err := conn.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	collection := conn.Database("test").Collection("sessions")
	filter := bson.D{{"_id", sess_id}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sess := &Session{}
	err = collection.FindOne(ctx, filter).Decode(sess)
	if err != nil {
		return nil, err
	} else {
		if IsExpiredSession(sess) {
			return nil, errors.New("Session is expired")
		} else {
			return sess, nil
		}
	}
}

func New(r *http.Request, w http.ResponseWriter, data map[string]interface{}) *Session {
	sess := newSession(data)
	// Cookie書き込み
	cookie := &http.Cookie{
		Name:    CookieName,
		Value:   sess.ID,
		Expires: sess.ExpireDt,
	}
	http.SetCookie(w, cookie)
	return sess
}

func Delete(r *http.Request, w http.ResponseWriter) error {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return err
	}
	sess_id := cookie.Value
	cookie.MaxAge = -1

	conn := models.GetConnection()
	defer func() {
		if err := conn.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := conn.Database("test").Collection("sessions")
	filter := bson.D{{"_id", sess_id}}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	http.SetCookie(w, cookie)
	return nil
}

func MakeSessionID(digit uint64) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-.?*^@:"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
