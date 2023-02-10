namespace go videoservice

// 用户结构体
struct User {
    1: i64 id // 用户唯一标识
    2: string name // 用户名
    3: i64 follow_count // 关注数
    4: i64 follower_count // 被关注数
    5: bool is_follow // 是否关注
}

// 视频结构体
struct Video {
    1: i64 id // 视频唯一标识
    2: User author // 发布者
    3: string play_url // 视频链接
    4: string cover_url // 封面地址
    5: i64 favorite_count // 视频的点赞总数
    6: i64 comment_count // 视频的评论总数
    7: bool is_favorite // true-已点赞，false-未点赞
    8: string title // 视频标题
}

// 评论结构体
struct Comment {
    1: i64 id // 视频评论id
    2: User user // 评论用户信息
    3: string content // 评论内容
    4: string create_date // 评论发布日期，格式 mm-dd
}

// 视频请求
struct FeedReq {
    1: optional string latest_time
    2: optional string token
}

// 视频响应
struct FeedResp {
    1: i64 status_code
    2: string status_msg
    3: i64 next_time
    4: list<Video> video_list
}

// 视频投稿接口请求
struct PublishActionReq {
    1: list<byte> data
    2: string token
    3: string title
}

// 视频投稿接口响应
struct PublishActionResp {
    1: i64 status_code
    2: string status_msg
}


// 评论操作接口请求
struct CommentActionReq {
    1: string token
    2: string video_id
    3: string action_type
    4: optional string comment_text
    5: optional string comment_id
}

// 评论操作接口响应
struct CommentActionResp {
    1: i64 status_code
    2: string status_msg
    3: Comment comment
}

// 关系操作接口请求
struct RelationActionReq {
    1: string token
    2: string to_user_id
    3: string action_type
}

// 关系操作接口响应
struct RelationActionResp {
    1: i64 status_code
    2: string status_msg

}

service VideoService {
    FeedResp Feed(1: FeedReq req) // 视频流接口
    PublishActionResp PublishAction(1: PublishActionReq req) // 视频投稿接口
    CommentActionResp CommentAction(1: CommentActionReq req) // 评论操作接口
    RelationActionResp RelationAction(1: RelationActionReq req) (api.post="/douyin/relation/action") // 关系操作接口
}