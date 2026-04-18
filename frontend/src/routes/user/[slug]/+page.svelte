<script lang="ts">
	let { data } = $props();
</script>

<svelte:head>
	<title>ZenQuest</title>
</svelte:head>

{#if data.user}
	<p>{data.user.email}</p>
	<!-- THIS IS FOR TESTING!! NEED TO MAKE A REAL USER PAGE -->
{/if}
