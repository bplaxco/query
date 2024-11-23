# Query

Look up resources on the internet.

## Status

Just getting started. Can do basic query file parsing.

## Example

What a query might look like

```
fetch
    source="google.mail.users.messages.list"
    userId="me" 
    q="subject:(Bill Available) 'Total'"
    maxResults=1 
| jq filter="{\"email_id\": .messages[0].id}" 
| fetch source=google.mail.users.messages.get userId="me" id=.email_id
| regex match="\$\d+\.\d+" name="total"
| jq raw_output=true filter=".total"
```

What the result could look like

```
$ query < last-bill.query
$123.45
```


## Syntax

```
cmd_list

cmd_list: [cmd [| cmd_list]]
cmd: name arg_list
name: string
arg_list: [key=value [arg_list]]
key: string
value: a jq filter applied to the current context 
```

## Commands

None of the commands are implemented yet

## Goal

Something a little like Splunk but cli based and using ad-hoc services as the
sources.
