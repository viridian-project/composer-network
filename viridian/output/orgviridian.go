package main
import "time"
type UserContact struct {
   Id string `json:"id"`
   Email string `json:"email"`
   Timestamp time.Time `json:"timestamp"`
   Verified bool `json:"verified"`
}
type UserSecret struct {
   Id string `json:"id"`
   Secret string `json:"secret"`
}
type User struct {
   Name string `json:"name"`
   CreatedAt time.Time `json:"createdAt"`
   Reputation int32 `json:"reputation"`
   LastCommentAt time.Time `json:"lastCommentAt"`
   LastCommentDeletedAt time.Time `json:"lastCommentDeletedAt"`
   AvatarUrl string `json:"avatarUrl"`
   PublicEmail string `json:"publicEmail"`
   Bio string `json:"bio"`
}
type Person struct {
   User
   Country string `json:"country"`
   PassportNrHash string `json:"passportNrHash"`
   RealName string `json:"realName"`
   Url string `json:"url"`
   Location string `json:"location"`
}
type Organization struct {
   User
   OrgName string `json:"orgName"`
   Url string `json:"url"`
   Country string `json:"country"`
   Address string `json:"address"`
}
type Company struct {
   Organization
}
type Status int
const (
   PRELIMINARY Status = 1 + iota
   ACTIVE
   OUTDATED
   DELETED
   REJECTED
)
type ReviewableAsset struct {
   Id string `json:"id"`
   CreatedAt time.Time `json:"createdAt"`
   Status Status `json:"status"`
}
type UpdatableAsset struct {
   ReviewableAsset
   UpdatedAt time.Time `json:"updatedAt"`
   ChangeReason string `json:"changeReason"`
}
type ScorableAsset struct {
   UpdatableAsset
   Score Score `json:"score"`
}
type Label struct {
   ScorableAsset
   Locales []LabelLocaleData `json:"locales"`
   Version string `json:"version"`
}
type Producer struct {
   ScorableAsset
   Name string `json:"name"`
   Address string `json:"address"`
   Url string `json:"url"`
}
type Product struct {
   ScorableAsset
   Gtin string `json:"gtin"`
   Locales []ProductLocaleData `json:"locales"`
}
type InfoCategory int
const (
   GENERAL_INFORMATION InfoCategory = 1 + iota
   LIFE_CYCLE_ANALYSIS
   EXTERNAL_COSTS
   STUDY_OR_PAPER
   PRESS_ARTICLE
   INVESTIGATIVE_REPORT
   CORPORATE_SOCIAL_RESPONSIBILITY
   JURISDICTION
   OTHER
)
type Information struct {
   UpdatableAsset
   Title string `json:"title"`
   Category InfoCategory `json:"category"`
   Description string `json:"description"`
   Sources []Source `json:"sources"`
   Weight int32 `json:"weight"`
}
type ReviewDecision int
const (
   PENDING ReviewDecision = 1 + iota
   APPROVED
   REJECTED
   IGNORED
)
type RejectReason int
const (
   INAPPROPRIATE RejectReason = 1 + iota
   INCORRECT
   OUTDATED
   DUPLICATE
   MISSING_SRC
   OTHER
)
type Review struct {
   Id string `json:"id"`
   RequestedAt time.Time `json:"requestedAt"`
   Decision ReviewDecision `json:"decision"`
   Timestamp time.Time `json:"timestamp"`
   RejectReason RejectReason `json:"rejectReason"`
   ReasonComment string `json:"reasonComment"`
}
type Rating struct {
   Id string `json:"id"`
   CreatedAt time.Time `json:"createdAt"`
   Score Score `json:"score"`
   Weight int32 `json:"weight"`
   Status Status `json:"status"`
}
type FlagReason int
const (
   INAPPROPRIATE FlagReason = 1 + iota
   INCORRECT
   OUTDATED
   TRIVIAL
   OTHER
)
type Comment struct {
   ReviewableAsset
   Text string `json:"text"`
   Weight int32 `json:"weight"`
   Flag Flag `json:"flag"`
}
type Voting struct {
   Id string `json:"id"`
   Timestamp time.Time `json:"timestamp"`
   Vote int32 `json:"vote"`
}
type RateVoting struct {
   Voting
}
type InfoVoting struct {
   Voting
}
type CommentVoting struct {
   Voting
}
