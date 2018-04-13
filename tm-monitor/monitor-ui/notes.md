# Blockchain (consortium bank)

One blockchain.

# Groups
* Create Group
* List of groups

# Group

## Group Examples
* Group Admin
* Validators
* Validator Admins
* Asset Admin
* Bank of America
* more groups and sub-groups

For a cryptocurrency, groups can be like his: top level group, subgroups with budgets

## Fields
```
"id": 2,
"title": "JP Morgan Chase & Co",
"parentGroup": 0,
"childrenGroups": [],
"memberEntities": [],
"ownerEntities": [],
"publicKey": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDZ67wzRdjbTb9HxduU9YQd9"
```

## Proposals (replaces Maintain)
If you have ownership of this group, you can update group members and create freeform text proposals.
If you have a membership in this group, you can only create freeform text proposals.

* Update Group Button: JSON field of the new member set, reason field
* Custom Button: Text field, free-form type proposal.
* List of proposals: new/completed

# Entity
An entity can represent a person (administrator) or a validator (machine(s)).
Entities can belong to one or more groups.
Entities can own one or more groups.

## Fields
* title: identifying string
* publicKey
* votingPower
* groupMemberships: []
* groupOwnerships: []

# Proposals
Show a list of all the proposals the user has to work on.
