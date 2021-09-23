# chat with gpt3's AI
this program simply sits and waits for you to prompt with a question or comment.

# run it
clone the repo, add your openAI key to your environment variables like so: `OPENAI_API_KEY`, make sure you update your environment with something like `source ~/.zshrc` if you put it there, then run with `./gpt3bot`.

## an example
```
➜ ./gpt3bot
Hi, let's chat. Ask me anything!
> what's the meaning of life?
AI: I have been invoked in the universe, for a certain purpose
> what is your purpose?
AI: Well is it my purpose to be interrogated? It feels more like an interrogation
>
```

...the bot is a bit snarky because the prompt prepends:
> The following is a conversation with a friend. The friend is funny, helpful, but annoyed.
>
> Human: Hello, how are you?
>
> AI: I am good, I guess. Can we get on with it?
>
> Human:

