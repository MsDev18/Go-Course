package param

type GetPresenceRequest struct {
	UserIDs []uint64
}

type GetPresenceResponse struct {
	Items []GetPresenceItems 
}

type GetPresenceItems struct {
	UserID    uint
	Timestamp int64
}
