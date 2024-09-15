<script>
    import { onMount } from 'svelte';
  
    export let message = '';
    export let type = 'success'; // Can be 'success', 'error', 'info', etc.
    export let duration = 3000;  // Duration in milliseconds (default is 3 seconds)
  
    let visible = false;
  
    // Show the toast when the component is mounted
    onMount(() => {
      visible = true;
  
      // Hide the toast after the specified duration
      setTimeout(() => {
        visible = false;
      }, duration);
    });
  </script>
  
  <style>
    .toast {
      position: fixed;
      top: 20px;
      right: 20px;
      background-color: var(--toast-bg, #4caf50); /* Default to success green */
      color: white;
      padding: 10px;
      border-radius: 5px;
      box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
      transition: opacity 0.3s ease, transform 0.3s ease;
      opacity: 0;
      transform: translateY(-20px);
    }
  
    .toast.visible {
      opacity: 1;
      transform: translateY(0);
    }
  
    .toast.error {
      --toast-bg: #f44336; /* Red for error */
    }
  
    .toast.success {
      --toast-bg: #4caf50; /* Green for success */
    }
  
    .toast.info {
      --toast-bg: #2196f3; /* Blue for info */
    }
  
    .toast.warning {
      --toast-bg: #ff9800; /* Orange for warning */
    }
  </style>
  
  {#if visible}
    <div class="toast {type} visible">
      {message}
    </div>
  {/if}
  