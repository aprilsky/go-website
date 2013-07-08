package cms;

import (
    //"log"
    "time"
    "strings"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

const BLOG_POST_COLL_NAME = "blog_posts"
type BlogPost struct {
    Id bson.ObjectId `bson:"_id,omitempty"`
    Url string
    Title string
    Content string
    Published bool
    PubDate time.Time //bson.MongoTimestamp
    Author string
    Tags []string
}

// Returns a list of blog post instances
func GetRecentBlogPosts(db *mgo.Database) ([]BlogPost, error) {
    var blogPosts []BlogPost
    var blogPostColl *mgo.Collection

    // Auto Disptach info objects
    blogPostColl = db.C(BLOG_POST_COLL_NAME)
    query := blogPostColl.Find(bson.M{"published":true}).Sort("-pubdate")

    err := query.All(&blogPosts)
    return blogPosts, err
}

// Returns the Id and the error
func InsertNewBlogPost(db *mgo.Database, post *BlogPost) error {
    var blogPostColl *mgo.Collection
    blogPostColl = db.C(BLOG_POST_COLL_NAME)

    // Default empty fields
    post.Id = bson.NewObjectId()
    post.PubDate = time.Now()

    // Insert
    err := blogPostColl.Insert(post)

    return err
}

// Loads and return a blog post from database
func GetBlogPost(db *mgo.Database, postId string) (BlogPost,error) {
    var blogPostColl *mgo.Collection
    blogPostColl = db.C(BLOG_POST_COLL_NAME)

    blogPost := BlogPost{}
    err := blogPostColl.Find(bson.M{"_id":bson.ObjectIdHex(postId)}).One(&blogPost)

    return blogPost, err
}

// Loads and return a blog post from database
func DeleteBlogPost(db *mgo.Database, postId string) error {
    var blogPostColl *mgo.Collection
    blogPostColl = db.C(BLOG_POST_COLL_NAME)

    return blogPostColl.Remove(bson.M{"_id":bson.ObjectIdHex(postId)})
}

// Returns the string in small letters, replacing spaces for hifens
func Slugify(s string) string {
    // TODO: clean letters to NFKD: http://play.golang.org/p/D7hmrTwi-i
    for strings.Contains(s, "  ") {
        s = strings.Replace(s, "  ", " ", -1)
    }
    s = strings.Replace(s, " ", "-", -1)
    return strings.ToLower(s)
}
