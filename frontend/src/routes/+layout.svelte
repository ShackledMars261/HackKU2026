<script lang="ts">
  import '../app.css';
  import { Spinner } from "$lib/components/ui/spinner/index";
  let { children } = $props();
</script>

<main class="relative flex flex-col h-screen overflow-hidden">
  <div class="relative flex overflow-hidden grow">
    {@render children()}
	<Spinner />
  </div>
</main>