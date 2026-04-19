<script lang="ts">
	import '../app.css';
	import NavBar from '@/components/NavBar.svelte';
	import favicon from '$lib/assets/favicon.svg';

	let { data, children } = $props();
</script>

<svelte:head>
	<link rel="icon" type="image/png" href={favicon} />
</svelte:head>

<main class="relative flex h-screen flex-col overflow-hidden">
	<div class="w-full shrink border-b border-primary/20 bg-accent/50 px-2 py-2">
		<NavBar isSignedIn={data.isSignedIn || false} />
	</div>
	<div class="relative flex grow overflow-hidden">
		{@render children()}
	</div>
</main>
