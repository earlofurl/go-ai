<script lang="ts">
    import {SendMessage} from '../wailsjs/go/main/App.js';

    type ChatCompletionMessage = {
        role: string
        content: string
        name?: string
    }

    let messages: Array<ChatCompletionMessage> = [];
    let userInput: string = "";
    let systemInput: string = "You are a helpful assistant.";
    let loading: boolean = false;

    async function sendMessage(): Promise<void> {
        if (!userInput) {
            return;
        }

        if (messages.length === 0) {
            messages = [
                {
                    role: "system",
                    content: systemInput,
                },
            ];
        }

        messages = [
            ...messages,
            {
                role: "user",
                content: userInput,
            },
        ];

        loading = true;
        try {
            const result = await SendMessage(messages);
            messages = [
                ...messages,
                {
                    role: "assistant",
                    content: result,
                },
            ];
        } catch (error) {
            messages = [
                ...messages,
                {
                    role: "assistant",
                    content: "Error: " + error.message,
                },
            ];
        } finally {
            userInput = "";
            loading = false;
        }
    }

    $: messages = [...messages];
</script>


<main>
    <div class="title">Wails AI Chat Assistant</div>
    <div class="instructions">Type a message and press send to get a response from the AI assistant.</div>
    <div class="input-box" id="input">
        <label for="system-message">System Message</label>
        <textarea autocomplete="off" bind:value={systemInput} class="input" id="system-message" rows="2"></textarea>
        <label for="user-message">User Message</label>
        <textarea autocomplete="off" bind:value={userInput} class="input" id="user-message" rows="4"></textarea>
        <button class="btn" on:click={() => sendMessage()} disabled={loading}>Send</button>
        {#if loading}
            <div class="spinner"></div>
        {/if}
    </div>
    <div class="response" id="response">
        {#each messages as message, index (index)}
            <div class="{message.role === 'user' ? 'user-message' : 'assistant-message'}">
                {message.role}: {message.content}
            </div>
        {/each}
    </div>
</main>

<style>
    body {
        font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
        margin: 0;
        padding: 0;
    }

    main {
        width: 80%;
        margin: 0 auto;
    }

    .title {
        text-align: center;
        font-size: 2rem;
        margin-bottom: 0.5rem;
    }

    .instructions {
        text-align: center;
        font-size: 1rem;
        margin-bottom: 1rem;
        color: #666;
    }

    .user-message {
        border-radius: 3px;
        padding: 10px;
        margin-bottom: 10px;
        width: 100%;
        color: #90EE90;
    }

    .assistant-message {
        border-radius: 3px;
        padding: 10px;
        margin-bottom: 10px;
        width: 100%;
        color: #fff;
    }

    .input-box {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .input-box label {
        font-size: 1.2rem;
        margin-bottom: 0.5rem;
    }

    .input-box .btn {
        width: 100px;
        height: 40px;
        line-height: 40px;
        border-radius: 3px;
        border: none;
        margin: 1rem 0 0;
        padding: 0 8px;
        cursor: pointer;
    }

    .input-box .btn:hover {
        background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
        color: #333333;
    }

    .input-box .input {
        border: none;
        border-radius: 3px;
        outline: none;
        padding: 10px;
        background-color: rgba(240, 240, 240, 1);
        -webkit-font-smoothing: antialiased;
        width: 100%;
        resize: none;
    }

    .input-box .input:hover {
        border: none;
        background-color: rgba(255, 255, 255, 1);
    }

    .input-box .input:focus {
        border: none;
        background-color: rgba(255, 255, 255, 1);
    }

    .response {
        margin-top: 1rem;
    }

    .spinner {
        border: 4px solid rgba(0, 0, 0, 0.1);
        width: 36px;
        height: 36px;
        border-radius: 50%;
        border-left-color: #09f;
        animation: spinner 1s linear infinite;
        margin-left: 10px;
    }

    .spinner {
        border: 4px solid rgba(0, 0, 0, 0.1);
        width: 36px;
        height: 36px;
        border-radius: 50%;
        border-left-color: #09f;
        animation: spinner 1s linear infinite;
    }

    @keyframes spinner {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }

</style>
