<script lang="ts">
	let { data } = $props();
</script>

<svelte:head>
	<title>SpotDrop</title>
</svelte:head>

{#if data.user}
	<p>{data.user.email}</p>
	<!-- THIS IS FOR TESTING!! NEED TO MAKE A REAL USER PAGE -->
{/if}
