# Open questions:

 * Are we happy with the edit conflict resolution? (While asset is under review,
   no other change can be proposed, it is locked. When a user starts editing
   while another user already edits and submits first, the second user is notified
   upon submit that a change is under review and the differences between the two
   users' versions are indicated.)
 * OK to not keep old versions of comments? (User who commented is allowed to edit
   comment, but no-one else. If comment is subsequently changed, no-one knows about
   it. So a flag might become obsolete, causing confusion.)
