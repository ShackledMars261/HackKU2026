<script lang="ts">
	import '../app.css';
	import { Spinner } from '$lib/components/ui/spinner/index';
	import NavBar from '@/components/NavBar.svelte';
	let { children } = $props();
</script>

<main class="relative flex h-screen flex-col overflow-hidden">
	<div class="shrink px-2 py-2">
		<NavBar />
	</div>
	<div class="relative flex grow overflow-hidden">
		{@render children()}
		<Spinner />
	</div>
</main>
