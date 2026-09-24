type Tweet struct {
	id int
	time int
}

type Twitter struct {
	time int
    tweets map[int][]Tweet
	following map[int]map[int]struct{}
}


func Constructor() Twitter {
	return Twitter{
		time : 0,
		tweets: make(map[int][]Tweet),
		following : make(map[int]map[int]struct{}),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.tweets[userId] = append(this.tweets[userId], Tweet{id: tweetId, time: this.time})
	this.time++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
	result := []Tweet{}
	tweets := this.tweets[userId]

	result = append(result, tweets...)
	for followerId := range this.following[userId] {
		result = append(result, this.tweets[followerId]...)
	}
    
	sort.Slice(result, func(i, j int) bool{
		return result[i].time > result[j].time
	})
	finalResult := []int{}
	for i := 0; i < 10 && i < len(result); i++ {
		finalResult = append(finalResult, result[i].id)
	}
	return finalResult
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
	if _, ok := this.following[followerId]; !ok {
		this.following[followerId] = make(map[int]struct{}) 
	}
	this.following[followerId][followeeId] = struct{}{}
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	delete(this.following[followerId], followeeId)
}
