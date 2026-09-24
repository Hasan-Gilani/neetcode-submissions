type Tweet struct {
    id   int
    time int
}

type Twitter struct {
    time      int
    tweets    map[int][]Tweet // userId -> their tweets, newest last
    following map[int]map[int]bool
}

func Constructor() Twitter {
    return Twitter{
        tweets:    make(map[int][]Tweet),
        following: make(map[int]map[int]bool),
    }
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
    t.tweets[userId] = append(t.tweets[userId], Tweet{id: tweetId, time: t.time})
    t.time++
}

func (t *Twitter) GetNewsFeed(userId int) []int {
    candidates := []Tweet{}

    // gather own tweets
    candidates = append(candidates, t.tweets[userId]...)

    // gather followees' tweets
    for followeeId := range t.following[userId] {
        candidates = append(candidates, t.tweets[followeeId]...)
    }

    // sort by time descending (most recent first)
    sort.Slice(candidates, func(i, j int) bool {
        return candidates[i].time > candidates[j].time
    })

    result := []int{}
    for i := 0; i < len(candidates) && i < 10; i++ {
        result = append(result, candidates[i].id)
    }
    return result
}

func (t *Twitter) Follow(followerId int, followeeId int) {
    if t.following[followerId] == nil {
        t.following[followerId] = make(map[int]bool)
    }
    t.following[followerId][followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
    delete(t.following[followerId], followeeId)
}