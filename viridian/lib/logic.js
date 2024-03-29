/*
 * Licensed under the GNU General Public License, Version 3 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.gnu.org/licenses/
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

'use strict';

// Our namespace:
const NS = 'org.viridian';

// For generating unique IDs
// import uuidv4 from 'uuid/v4';

/**
 * Let a personal user update their own data with this function
 * @param {org.viridian.UpdatePersonData} tx - The transaction instance
 * @transaction
 */
async function updatePersonData(tx) {
  const personRegistry = await getParticipantRegistry(NS + '.Person');
  if (tx.avatarUrl) { // if something was entered
    tx.person.avatarUrl = tx.avatarUrl;
  }
  if (tx.publicEmail) { // if something was entered
    tx.person.publicEmail = tx.publicEmail;
  }
  if (tx.bio) { // if something was entered
    tx.person.bio = tx.bio;
  }
  if (tx.realName) { // if something was entered
    tx.person.realName = tx.realName;
  }
  if (tx.url) { // if something was entered
    tx.person.url = tx.url;
  }
  if (tx.location) { // if something was entered
    tx.person.location = tx.location;
  }
  // may throw an error with:
  // throw new Error(`Something wrong when trying to change Person ${tx.person}.`);
  await personRegistry.update(tx.person);
}

/**
 * Let an organizational user update their own data with this function
 * @param {org.viridian.UpdateOrganizationData} tx - The transaction instance
 * @transaction
 */
async function updateOrganizationData(tx) {
  const organizationRegistry = await getParticipantRegistry(NS + '.Organization');
  if (tx.avatarUrl) { // if something was entered
    tx.organization.avatarUrl = tx.avatarUrl;
  }
  if (tx.publicEmail) { // if something was entered
    tx.organization.publicEmail = tx.publicEmail;
  }
  if (tx.bio) { // if something was entered
    tx.organization.bio = tx.bio;
  }
  if (tx.orgName) { // if something was entered
    tx.organization.orgName = tx.orgName;
  }
  if (tx.url) { // if something was entered
    tx.organization.url = tx.url;
  }
  if (tx.country) { // if something was entered
    tx.organization.country = tx.country;
  }
  if (tx.address) { // if something was entered
    tx.organization.address = tx.address;
  }
  // may throw an error with:
  // throw new Error(`Something wrong when trying to change Organization ${tx.organization}.`);
  await organizationRegistry.update(tx.organization);
}

/**
 * Let a user update their contact data with this function
 * @param {org.viridian.UpdateUserContact} tx - The transaction instance
 * @transaction
 */
async function updateUserContact(tx) {
  const factory = getFactory();

  // Create a new user contact with the new data
  // var id = uuidv4();
  var id = tx.contactId;
  const contact = factory.newResource(NS, 'UserContact', id);
  contact.user = factory.newRelationship(NS, 'User', tx.user.getIdentifier());
  contact.email = tx.email;
  contact.timestamp = new Date();

  // const userRegistry = await getParticipantRegistry(NS + '.User');
  const contactRegistry = await getAssetRegistry(contact.getFullyQualifiedType());
  await contactRegistry.add(contact);

  // Create a new user secret that will be used to verify the new user contact
  // id = uuidv4();
  id = tx.secretId;
  const secret = factory.newResource(NS, 'UserSecret', id);
  secret.contact = factory.newRelationship(NS, 'UserContact', contact.getIdentifier());
  // secret.secret = uuidv4();
  secret.secret = '0ec40608-a4d1-4ae7-8711-fd6e292d4bc6'; // hardcoded for now, must get this
    // random number from an external resource so that it is deterministic (equal across all peers)

  // Send an email to the new email addres with the secret:
  // TODO
  // Email also must be sent by external resource, so that only one email is sent and not
  // as many as there are peers

  // Store the secret for verification later:
  const secretRegistry = await getAssetRegistry(NS + '.UserSecret');
  await secretRegistry.add(secret);
}
