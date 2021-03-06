package main

import (
    "strings"

    "github.com/nlopes/slack"
)

// REMOVE FRIEND
func remove_friend(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) != 2 {
        send_message("Usage: *addfriend <name>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])
    yaml := read_yaml(YAML_FILE)

    // check if friend exists
    _, exists := yaml.Friends[name]

    if exists {

        // remove friend and write
        delete(yaml.Friends, name)
        write_yaml(YAML_FILE, yaml)

    }

}

// REMOVE PHRASE
func remove_phrase(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 3 {
        send_message("Usage: *addphrase <name> <phrase>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])
    phrase := strings.ToLower(strings.Join(tokens[2:], " "))
    yaml := read_yaml(YAML_FILE)

    _, exists := yaml.Friends[name]

    if exists {

        // remove phrase and write
        friend := yaml.Friends[name]
        friend.Phrases = delete_item(friend.Phrases, phrase)
        yaml.Friends[name] = friend
        write_yaml(YAML_FILE, yaml)
    }

}

// REMOVE JOKE
func remove_joke(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 2 {
        send_message("Usage: *addjoke <joke>", ev, rtm)
    }

    joke := strings.ToLower(strings.Join(tokens[1:], " "))
    yaml := read_yaml(YAML_FILE)

    // remove joke and write
    yaml.Jokes = delete_item(yaml.Jokes, joke)
    write_yaml(YAML_FILE, yaml)

}

// REMOVE INSULT
func remove_insult(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 2 {
        send_message("Usage: *addinsult <insult>", ev, rtm)
    }

    insult := strings.ToLower(strings.Join(tokens[1:], " "))
    yaml := read_yaml(YAML_FILE)

    // remove insult and write
    yaml.Insults = delete_item(yaml.Insults, insult)
    write_yaml(YAML_FILE, yaml)

}

// REMOVE ALIAS
func remove_alias(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 3 {
        send_message("Usage: *addalias <name> <alias>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])
    alias := strings.ToLower(strings.Join(tokens[2:], " "))
    yaml := read_yaml(YAML_FILE)

    _, exists := yaml.Friends[name]

    if exists == true {

        // remove alias and write
        friend := yaml.Friends[name]
        friend.Aliases = delete_item(friend.Aliases, alias)
        yaml.Friends[name] = friend
        write_yaml(YAML_FILE, yaml)

    }

}
