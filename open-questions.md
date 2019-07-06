# Open questions:

 * Identify users by their user name? User name must be unique? Could also create
   a unique ID for each user (e.g. on StackOverflow, it's an integer number), so
   that the user name can be changed at will and must not be unique. But github
   does not follow this path: User name cannot be changed and probably must be
   unique.
 * What if user wants to change their user name? Users are identified by their
   user name, so it may be a problem to change that name. They also cannot open
   a new account under the desired name if their passport is used to ensure that
   every person only has one account.
 * Are we happy with the edit conflict resolution? (While asset is under review,
   no other change can be proposed, it is locked. When a user starts editing
   while another user already edits and submits first, the second user is notified
   upon submit that a change is under review and the differences between the two
   users' versions are indicated.)
 * OK to not keep old versions of comments? (User who commented is allowed to edit
   comment, but no-one else. If comment is subsequently changed, no-one knows about
   it. So a flag might become obsolete, causing confusion.)
 * Is it OK that the time of last comment and comment deletion are stored in the user
   asset? The user should not be able to edit this DateTime at will, because it would
   circumvent the comment blocking feature.
 * Are the DateTimes for time of last comment and comment deletion even necessary?
   Maybe it's efficient enough to calculate if from a query each time the user wants
   to add another comment?
 * Are there too many reviews for every addition and change in the assets? Is it
   simply too much work for the users?
 * Should we progressively loosen restrictions on users based on reputation? For
   example, StackOverflow uses similar peer reviews, but users with reputation
   above 2000 do not require peer reviews of their edits. (Does SO have no peer
   reviews for added questions/answers?)
 * Do we need more modeling, e.g. assets, for the review process? The review
   request should be withdrawn from users if they do not respond within a
   reasonable time. Maybe one could also select users for review not entirely
   random, but with a bias towards more active users, in order to increase
   review productivity.
