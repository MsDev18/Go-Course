package param

type GetPresenceRequest struct {
	UserIDs []uint
}

type GetPresenceResponse struct {
	Items []GetPresenceItems 
}

type GetPresenceItems struct {
	UserID    uint
	Timestamp int64
}
