package main

import "fmt"

type Tweet struct {
	ID   int
	User *User
	Next *Tweet
}

type User struct {
	ID        int
	Followers map[int]int
}

type Twitter struct {
	Tweets *Tweet
	Users  map[int]*User
}

func Constructor() Twitter {
	return Twitter{
		Users: map[int]*User{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.Users[userId]; !ok {
		this.Users[userId] = &User{ID: userId, Followers: map[int]int{}}
	}
	user := this.Users[userId]
	newTweet := &Tweet{ID: tweetId, User: user}
	if this.Tweets == nil {
		this.Tweets = newTweet
	} else {
		next := this.Tweets
		this.Tweets = newTweet
		newTweet.Next = next
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	tweets := []int{}
	followees := map[int]int{userId: userId}

	for _, user := range this.Users {
		if _, ok := user.Followers[userId]; ok {
			followees[user.ID] = user.ID
		}
	}
	curr := this.Tweets
	for curr != nil {
		if _, ok := followees[curr.User.ID]; ok {
			tweets = append(tweets, curr.ID)
		}
		if len(tweets) == 10 {
			return tweets
		}
		curr = curr.Next
	}
	return tweets
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.Users[followerId]; !ok {
		this.Users[followerId] = &User{ID: followerId, Followers: make(map[int]int)}
	}

	if _, ok := this.Users[followeeId]; !ok {
		this.Users[followeeId] = &User{ID: followeeId, Followers: make(map[int]int)}
	}

	this.Users[followeeId].Followers[followerId] = followerId
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	user := this.Users[followeeId]

	if _, ok := user.Followers[followerId]; ok {
		delete(user.Followers, followerId)
	}
}

func main() {
	obj := Constructor()
	obj.PostTweet(1, 6765) //
	obj.PostTweet(5, 671)
	obj.PostTweet(3, 2868) //
	obj.PostTweet(4, 8148) //
	obj.PostTweet(4, 386)  //
	obj.PostTweet(3, 6673) //
	obj.PostTweet(3, 7946) //
	obj.PostTweet(3, 1445) //
	obj.PostTweet(4, 4822) //
	obj.PostTweet(1, 3781) //
	obj.PostTweet(4, 9038) //
	obj.PostTweet(1, 9643) //
	obj.PostTweet(3, 5917) //
	obj.PostTweet(2, 8847)
	obj.Follow(1, 3)
	obj.Follow(1, 4)
	obj.Follow(4, 2)
	obj.Follow(4, 1)
	obj.Follow(3, 2)
	obj.Follow(3, 5)
	obj.Follow(3, 1)
	obj.Follow(2, 3)
	obj.Follow(2, 1)
	obj.Follow(2, 5)
	obj.Follow(5, 1)
	obj.Follow(5, 2)
	fmt.Println(obj.GetNewsFeed(1))
	fmt.Println(obj.GetNewsFeed(2))
	fmt.Println(obj.GetNewsFeed(3))
	fmt.Println(obj.GetNewsFeed(4))
	fmt.Println(obj.GetNewsFeed(5))

}
