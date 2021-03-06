package circleci

var (
	testProjectResponse     string
	testProjectRecentBuilds string
	testTriggerBuild        string
)

func init() {
	setupTestProjects()
	setupTestProjectRecentBuilds()
	setupTestTriggerBuild()
}

func setupTestProjects() {
	testProjectResponse = `{
  "irc_server": null,
  "scopes": [
    "write-settings",
    "view-builds",
    "read-settings",
    "trigger-builds",
    "all",
    "status",
    "none"
  ],
  "irc_keyword": null,
  "followed": true,
  "vcs-type": "github",
  "aws": {
    "keypair": null
  },
  "slack_webhook_url": null,
  "flowdock_api_token": null,
  "parallel": 1,
  "username": "mtchavez",
  "campfire_room": null,
  "extra": "",
  "branches": {
    "master": {
      "running_builds": [],
      "recent_builds": [
        {
          "build_num": 286,
          "status": "success",
          "outcome": "success",
          "added-at": "2016-10-01T04:16:40.186Z"
        },
        {
          "build_num": 285,
          "status": "success",
          "outcome": "success",
          "added-at": "2016-10-01T04:17:07.233Z"
        },
        {
          "build_num": 284,
          "status": "success",
          "outcome": "success",
          "added-at": "2016-10-01T03:38:39.317Z"
        },
        {
          "build_num": 283,
          "status": "success",
          "outcome": "success",
          "added-at": "2016-10-01T03:38:37.553Z"
        },
        {
          "build_num": 280,
          "status": "success",
          "outcome": "success",
          "added-at": "2016-09-25T01:36:03.375Z"
        }
      ],
      "pusher_logins": [
        "mtchavez"
      ]
    }
  },
  "jira": null,
  "slack_subdomain": null,
  "following": true,
  "setup": "",
  "campfire_subdomain": null,
  "slack_notify_prefs": null,
  "irc_password": null,
  "vcs_url": "https://github.com/mtchavez/circleci",
  "default_branch": "master",
  "hipchat_api_token": null,
  "irc_username": null,
  "language": "Ruby",
  "slack_channel_override": null,
  "hipchat_notify": null,
  "slack_api_token": null,
  "has_usable_key": true,
  "irc_notify_prefs": null,
  "campfire_token": null,
  "slack_channel": null,
  "feature_flags": {
    "trusty-beta": false,
    "osx": false,
    "set-github-status": true,
    "build-prs-only": false,
    "forks-receive-secret-env-vars": false,
    "fleet": null,
    "build-fork-prs": false,
    "autocancel-builds": false,
    "junit": true,
    "oss": true
  },
  "campfire_notify_prefs": null,
  "hipchat_room": null,
  "post_dependencies": "",
  "heroku_deploy_user": null,
  "irc_channel": null,
  "oss": true,
  "reponame": "circleci",
  "hipchat_notify_prefs": null,
  "compile": "",
  "dependencies": "",
  "ssh_keys": [
    {
      "hostname": "hostname",
      "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDC7Ca3FXm9nAC5kpIQ+uB2a7ShJkj4PvjNJSsI1rmz2Z8eXDMLOEW08qGXBjFWlDtZkGtAw2U8UmaHKPcJ3YM+Ss5sw0sRURW3pPpqF0JzKvDaNg8UxDLBk8CFxb2YLjjSpM2G5BMY5w38Af/U/8Hhb284lg7sdbtM9rRLlDeYVEx+T/nVOWsH+j3SAICZfpKb7GCz1OCbH2lhEXQqPMin+0UYw4ePrB0NvtC/x55fQ+cK1K3jbyVkZIXRw1KZAhUjEGCLu3CYsqkWCpLUsAABv36Vdb7KH9/kWXhUGKsS8LeiVLrJQDFaENDbvBo4rum1ORhgbD72DtBZKON4r6/",
      "fingerprint": "83:3d:57:76:98:96:f6:ae:20:57:18:84:a6:4d:58:9d"
    },
    {
      "hostname": "hostname",
      "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDC7Ca3FXm9nAC5kpIQ+uB2a7ShJkj4PvjNJSsI1rmz2Z8eXDMLOEW08qGXBjFWlDtZkGtAw2U8UmaHKPcJ3YM+Ss5sw0sRURW3pPpqF0JzKvDaNg8UxDLBk8CFxb2YLjjSpM2G5BMY5w38Af/U/8Hhb284lg7sdbtM9rRLlDeYVEx+T/nVOWsH+j3SAICZfpKb7GCz1OCbH2lhEXQqPMin+0UYw4ePrB0NvtC/x55fQ+cK1K3jbyVkZIXRw1KZAhUjEGCLu3CYsqkWCpLUsAABv36Vdb7KH9/kWXhUGKsS8LeiVLrJQDFaENDbvBo4rum1ORhgbD72DtBZKON4r6/",
      "fingerprint": "83:3d:57:76:98:96:f6:ae:20:57:18:84:a6:4d:58:9d"
    },
    {
      "hostname": "hostname",
      "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDC7Ca3FXm9nAC5kpIQ+uB2a7ShJkj4PvjNJSsI1rmz2Z8eXDMLOEW08qGXBjFWlDtZkGtAw2U8UmaHKPcJ3YM+Ss5sw0sRURW3pPpqF0JzKvDaNg8UxDLBk8CFxb2YLjjSpM2G5BMY5w38Af/U/8Hhb284lg7sdbtM9rRLlDeYVEx+T/nVOWsH+j3SAICZfpKb7GCz1OCbH2lhEXQqPMin+0UYw4ePrB0NvtC/x55fQ+cK1K3jbyVkZIXRw1KZAhUjEGCLu3CYsqkWCpLUsAABv36Vdb7KH9/kWXhUGKsS8LeiVLrJQDFaENDbvBo4rum1ORhgbD72DtBZKON4r6/",
      "fingerprint": "83:3d:57:76:98:96:f6:ae:20:57:18:84:a6:4d:58:9d"
    },
    {
      "hostname": "hostname",
      "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDC7Ca3FXm9nAC5kpIQ+uB2a7ShJkj4PvjNJSsI1rmz2Z8eXDMLOEW08qGXBjFWlDtZkGtAw2U8UmaHKPcJ3YM+Ss5sw0sRURW3pPpqF0JzKvDaNg8UxDLBk8CFxb2YLjjSpM2G5BMY5w38Af/U/8Hhb284lg7sdbtM9rRLlDeYVEx+T/nVOWsH+j3SAICZfpKb7GCz1OCbH2lhEXQqPMin+0UYw4ePrB0NvtC/x55fQ+cK1K3jbyVkZIXRw1KZAhUjEGCLu3CYsqkWCpLUsAABv36Vdb7KH9/kWXhUGKsS8LeiVLrJQDFaENDbvBo4rum1ORhgbD72DtBZKON4r6/",
      "fingerprint": "83:3d:57:76:98:96:f6:ae:20:57:18:84:a6:4d:58:9d"
    }
  ]
}`
}

func setupTestProjectRecentBuilds() {
	testProjectRecentBuilds = `[{
  "vcs_url" : "https://github.com/circleci/mongofinil",
  "build_url" : "https://circleci.com/gh/circleci/mongofinil/22",
  "build_num" : 22,
  "branch" : "master",
  "vcs_revision" : "1d231626ba1d2838e599c5c598d28e2306ad4e48",
  "committer_name" : "Allen Rohner",
  "committer_email" : "arohner@gmail.com",
  "subject" : "Don't explode when the system clock shifts backwards",
  "body" : "",
  "why" : "github",
  "dont_build" : null,
  "queued_at" : "2013-02-12T21:33:30Z",
  "start_time" : "2013-02-12T21:33:38Z",
  "stop_time" : "2013-02-12T21:34:01Z",
  "build_time_millis" : 23505,
  "username" : "circleci",
  "reponame" : "mongofinil",
  "lifecycle" : "finished",
  "outcome" : "failed",
  "status" : "failed",
  "retry_of" : null,
  "previous" : {
    "status" : "failed",
    "build_num" : 21
  }
}]`
}

func setupTestTriggerBuild() {
	testTriggerBuild = `
{
  "author_name": "Allen Rohner",
  "build_url": "https://circleci.com/gh/circleci/mongofinil/54",
  "reponame": "mongofinil",
  "failed": null,
  "infrastructure_fail": false,
  "canceled": false,
  "all_commit_details": [{
      "author_name": "Allen Rohner",
      "commit": "f1baeb913288519dd9a942499cef2873f5b1c2bf",
      "author_login": "arohner",
      "committer_login": "arohner",
      "committer_name": "Allen Rohner",
      "body": "Minor version bump",
      "author_date": "2014-04-17T08:41:40Z",
      "committer_date": "2014-04-17T08:41:40Z",
      "commit_url": "https://github.com/circleci/mongofinil/commit/f1baeb913288519dd9a942499cef2873f5b1c2bf",
      "committer_email": "arohner@gmail.com",
      "author_email": "arohner@gmail.com",
      "subject": "Merge pull request #15 from circleci/minor-version-bump"
    }],
  "previous": {
    "build_num": 53,
    "status": "success",
    "build_time_millis": 55413
  },
  "ssh_enabled": null,
  "author_email": "arohner@gmail.com",
  "why": "edit",
  "build_time_millis": null,
  "committer_email": "arohner@gmail.com",
  "parallel": 2,
  "retries": null,
  "compare": null,
  "dont_build": null,
  "committer_name": "Allen Rohner",
  "usage_queued_at": "2014-04-29T12:56:55.338Z",
  "branch": "master",
  "body": "Minor version bump",
  "author_date": "2014-04-17T08:41:40Z",
  "node": null,
  "committer_date": "2014-04-17T08:41:40Z",
  "start_time": null,
  "stop_time": null,
  "lifecycle": "not_running",
  "user": {
    "email": "arohner@gmail.com",
    "name": "Allen Rohner",
    "login": "arohner",
    "is_user": true
  },
  "subject": "Merge pull request #15 from circleci/minor-version-bump",
  "messages": [],
  "job_name": null,
  "retry_of": null,
  "previous_successful_build": {
    "build_num": 53,
    "status": "success",
    "build_time_millis": 55413
  },
  "outcome": null,
  "status": "not_running",
  "vcs_revision": "f1baeb913288519dd9a942499cef2873f5b1c2bf",
  "vcs_tag": "v0.1",
  "build_num": 54,
  "username": "circleci",
  "vcs_url": "https://github.com/circleci/mongofinil",
  "timedout": false
}
`
}
