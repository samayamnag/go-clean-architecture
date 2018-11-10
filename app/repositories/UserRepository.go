package repositories

import (
	"github.com/samayamnag/boilerplate/app/models/mongo"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

//Reader interface
type UserReaderInterface interface {
	Find(id mongo.ID) (*mongo.User, error)
	Search(query string) ([]*mongo.User, error)
	FindAll() ([]*mongo.User, error)
	FindByEmail(email string) (*mongo.User, error)
	Count() (int, error)
}

//UserWriterInterface user writer
type UserWriterInterface interface {
	Store(u mongo.User) (mongo.User, error)
	Update(u *mongo.User) (*mongo.User, error)
	Delete(id mongo.ID) error
}

//UserWriterInterface repository interface
type UserRepoInterface interface {
	UserReaderInterface
	UserWriterInterface
}

//UserUseCaseInterface use case interface
type UserUseCaseInterface interface {
	UserReaderInterface
	UserWriterInterface
}

//Store an user
func (r *MongoRepository) Count() (int, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C(USER_COLLECTION)
	total, err := coll.Find(nil).Count()
	switch err {
		case nil:
			return total, nil
		case mgo.ErrNotFound:
			return 0, mongo.ErrNotFound
		default:
			return 0, err
	}
}

//Store an user
func (r *MongoRepository) Store(u mongo.User) (mongo.User, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C(USER_COLLECTION)
	err := coll.Insert(u)
	if err != nil {
		return mongo.User{}, err
	}
	return u, nil
}

//Find user
func (r *MongoRepository) Find(id mongo.ID) (*mongo.User, error) {
	result := mongo.User{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C(USER_COLLECTION)
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, mongo.ErrNotFound
	default:
		return nil, err
	}
}

//Find user by email
func (r *MongoRepository) FindByEmail(email string) (*mongo.User, error) {
	result := mongo.User{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C(USER_COLLECTION)
	err := coll.Find(bson.M{"email": email}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, mongo.ErrNotFound
	default:
		return nil, err
	}
}

//FindAll users
func (r *MongoRepository) FindAll() ([]*mongo.User, error) {
	var d []*mongo.User
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C(USER_COLLECTION)
	err := coll.Find(nil).Sort("full_name").All(&d)
	switch err {
	case nil:
		return d, nil
	case mgo.ErrNotFound:
		return nil, mongo.ErrNotFound
	default:
		return nil, err
	}
}

//Search users
func (r *MongoRepository) Search(query string) ([]*mongo.User, error) {
	var d []*mongo.User
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C(USER_COLLECTION)
	conditioner := bson.M{"full_name": &bson.RegEx{Pattern: query, Options: "i"}}
	err := coll.Find(conditioner).Limit(10).Sort("full_name").All(&d)
	switch err {
	case nil:
		return d, nil
	case mgo.ErrNotFound:
		return nil, mongo.ErrNotFound
	default:
		return nil, err
	}
}

//Store an user
func (r *MongoRepository) Update(u *mongo.User) (*mongo.User, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C(USER_COLLECTION)
	err := coll.UpdateId(u.ID, u)
	if err != nil {
		return u, err
	}
	return u, nil
}

//Delete a user
func (r *MongoRepository) Delete(id mongo.ID) error {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C(USER_COLLECTION)
	return coll.Remove(bson.M{"_id": id})
}
