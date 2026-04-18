<script lang="ts">
	let { data } = $props();
</script>

<svelte:head>
	<title>ZenQuest</title>
</svelte:head>

{#if data.location}
	<p>{data.location.name}</p>
	<!-- THIS IS FOR TESTING!! NEED TO MAKE A REAL USER PAGE -->
{/if}
