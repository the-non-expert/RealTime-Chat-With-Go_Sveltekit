<!-- <script>
  import "../app.css";
</script>

<div class="">
  <div class=""></div>

  <div class="">
    <main>
      <slot />
    </main>
  </div>
</div> -->

<script lang="ts">
  import { onMount } from "svelte";

  let socket: WebSocket;
  let messages: string[] = [];
  let message = "";

  onMount(() => {
    // Replace with your backend WebSocket URL
    socket = new WebSocket("ws://localhost:8080/ws");

    socket.onopen = () => {
      console.log("Connected to WebSocket server");
    };

    socket.onmessage = (event: MessageEvent) => {
      messages = [...messages, event.data];
      console.log(messages);
      console.log("Received message from server: ", event.data);
    };

    socket.onclose = () => {
      console.log("Disconnected from WebSocket server");
    };

    socket.onerror = (error) => {
      console.error("WebSocket Error:", error);
    };

    return () => {
      socket.close();
    };
  });

  function sendMessage() {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(message);
      message = "";
    }
  }
</script>

<div class="chat-container">
  <div class="messages">
    {#each messages as msg}
      <div>{msg}</div>
    {/each}
  </div>
  <div class="message-input">
    <input bind:value={message} placeholder="Enter your message" />
    <button on:click={sendMessage}>Send</button>
  </div>
</div>

<style>
  /* Add some basic styling */
  .chat-container {
    max-width: 500px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
  }

  .messages {
    max-height: 300px;
    overflow-y: auto;
    border-bottom: 1px solid #ccc;
    margin-bottom: 10px;
    padding-bottom: 10px;
  }

  .message-input {
    display: flex;
  }

  .message-input input {
    flex: 1;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px 0 0 5px;
  }

  .message-input button {
    padding: 10px;
    border: 1px solid #ccc;
    border-left: 0;
    border-radius: 0 5px 5px 0;
    background-color: #007bff;
    color: white;
  }
</style>
