<script lang="ts">
	import { enhance } from '$app/forms';
	import Geolocation from 'svelte-geolocation';
	let coords: [number, number] = $state([0, 0]);
	let longitude: number = $derived(coords[0]);
	let latitude: number = $derived(coords[1]);
</script>

<Geolocation getPosition bind:coords />

<form method="POST" action="?/submitLocation" use:enhance>
	<input id="longitude" name="latitude" type="number" step="any" value={longitude} />
	<input id="latitude" name="latitude" type="number" step="any" value={latitude} />
	<label for="name">Location name:</label>
	<input id="name" name="name" type="text" />
	<button type="submit">Submit</button>
</form>
