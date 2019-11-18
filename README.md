# CS3083-Project
Project for Introduction to Databases

## PriCoSha
- Adam Sanghera
- Deryck Wong
- Gary Chi
- Minhazur Bhuiyan
- Ryan Rozbiani

## Features
- Login and Register (2)
- Manage posts (includes creating, editing, and deleting (which cascades deleting of shares, tags, comments)) (3)
- Create and delete tags (2)
- Create, delete, edit comments (3)
- Create, join/leave, delete friend groups (4)
- Personalize color (1)
- Change account name (1)
- Delete account (1)

## How to run

### Pre-Reqs

1.  Install node.js (on macs with homebrew, `brew install node`)
2.  Run `npm i -g serve`
3.  Install golang (on macs with homebrew, `brew install golang`)
4.  Run `go get -u adamsanghera/hashing`
5.  Run `go get -u adamsanghera/mysqlBus`
6.  Install mysql (on macs with homebrew, `brew install mysql`)

### Running it locally


0.  Make sure that ports 3000 and 5000 are open
1.  Start mysqld with default settings (root user has no password, on port 3306)
2.  Navigate to `/static-server` and run `serve` in one terminal tab, keep this open.
3.  Navigate to `/middleware` and run `sudo ./runGo.sh`, keep this open.
4.  Open `localhost:5000/index` in your browser.
5.  Enjoy!

## Scractchpad
When a person leaves a friend group, their tags in private posts that have been shared with that friend group are hidden.

When a person makes a public post, everyone can see it.

When a person makes a private post, only they can see it until they share it to friend groups.  Only the content creator can share their private post to friend groups.

Sharing does not have meaning for public posts.

In comments on private posts shared to a given friend group, maybe indicate whether the user exists

When a user leaves a group, give them the option to delete all of the traces of their presence in that group. Replaces username with '[deleted]'

Say what you mean, mean what you say.  Comments and posts are forever, even if you leave.  If you leave, name is anonymized.