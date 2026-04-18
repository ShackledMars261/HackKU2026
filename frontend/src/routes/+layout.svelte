<script lang="ts">
	import '../app.css';
	import NavBar from '@/components/NavBar.svelte';
	let { children } = $props();
</script>

<main class="relative flex h-screen flex-col overflow-hidden">
	<div class="shrink px-2 py-2 bg-accent/50 border-b border-primary/20">
		<NavBar />
	</div>
	<div class="relative flex grow overflow-hidden">
		{@render children()}
	</div>
</main>
