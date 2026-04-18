<script lang="ts">
	import { enhance } from '$app/forms';
	import Geolocation from 'svelte-geolocation';
	let coords: [number, number] = $state([0, 0]);
	let longitude: number = $derived(coords[0]);
	let latitude: number = $derived(coords[1]);
</script>

<Geolocation getPosition bind:coords />

<form method="POST" action="?/submitLocation" use:enhance>
	<input id="longitude" name="latitude" type="number" step="any" value={longitude} /><br />
	<input id="latitude" name="latitude" type="number" step="any" value={latitude} /><br />
	<label for="name">Location name:</label>
	<input id="name" name="name" type="text" /><br />
	<button type="submit">Submit</button><br />
</form>
<hr />
<br />
<form method="POST" action="?/submitPost" enctype="multipart/form-data">
	<br />
	<label for="location_id">Location id:</label>
	<input id="location_id" name="location_id" type="text" /><br />
	<label for="description">Description:</label>
	<input id="description" name="description" type="text" /><br />
	<label for="rating">Rating:</label>
	<input id="rating" name="rating" type="text" /><br />
	<input type="file" name="photos" accept="image/*" /><br />
	<button type="submit">Upload</button><br />
</form>
