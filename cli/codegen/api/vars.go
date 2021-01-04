package api



type User struct { 
}

// 获取唯一id
type SingleIDGetRequest struct { 
}


type SingleIDGetResponse struct { 
  	u *User `json:"user"`  
}

// 批量获取唯一id
// check
type BatchIDGetRequest struct { 
  	Count int32 `json:"count,string"` // 要获取的数量 
}


type BatchIDGetResponse struct { 
  	IDs []int64 `json:"ids"` // 
}
