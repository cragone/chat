# chat

Generate azure creds:

```az ad sp create-for-rbac --name "github-actions-sp" --role Contributor --scopes /subscriptions/YOUR_SUBSCRIPTION_ID

```

get azure login server

```az acr list --output table

```
