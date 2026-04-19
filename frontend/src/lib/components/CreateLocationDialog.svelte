<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';

	let open = $state(false);
	let name = $state('');
	let latitude = $state('');
	let longitude = $state('');
	let loading = $state(false);
	let error = $state<string | null>(null);

	async function locate() {
		await new Promise<void>((resolve, reject) => {
			navigator.geolocation.getCurrentPosition(
				(pos) => {
					latitude = pos.coords.latitude.toString();
					longitude = pos.coords.longitude.toString();
					resolve();
				},
				() => reject(new Error('Could not get location')),
				{ enableHighAccuracy: true, timeout: 10000 }
			);
		}).catch((e) => {
			error = e.message;
		});
	}

	async function handleSubmit() {
		error = null;

		if (!name.trim()) {
			error = 'Name is required.';
			return;
		}
		if (!latitude || !longitude) {
			error = 'Coordinates are required.';
			return;
		}

		loading = true;
		try {
			const resp = await fetch('/api/location', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					name: name.trim(),
					latitude: parseFloat(latitude),
					longitude: parseFloat(longitude)
				})
			});

			if (!resp.ok) throw new Error('Failed to create location.');

			open = false;
			name = '';
			latitude = '';
			longitude = '';
		} catch (e: any) {
			error = e.message;
		} finally {
			loading = false;
		}
	}
</script>

<Dialog.Root bind:open>
	<Dialog.Trigger>
		{#snippet child({ props })}
			<button
				{...props}
				class="inline-flex items-center justify-center gap-1 rounded-md px-4 py-2 text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground focus:outline-none"
			>
				Add Location
			</button>
		{/snippet}
	</Dialog.Trigger>

	<Dialog.Content class="sm:max-w-md">
		<Dialog.Header>
			<Dialog.Title>Add Location</Dialog.Title>
			<Dialog.Description>Enter the details for the new location.</Dialog.Description>
		</Dialog.Header>

		<div class="flex flex-col gap-4 py-2">
			<div class="flex flex-col gap-1.5">
				<Label for="name">Name</Label>
				<Input id="name" placeholder="e.g. Central Park" bind:value={name} />
			</div>

			<div class="flex flex-col gap-1.5">
				<Label>Coordinates</Label>
				<div class="flex gap-2">
					<Input placeholder="Latitude" bind:value={latitude} type="number" step="any" />
					<Input placeholder="Longitude" bind:value={longitude} type="number" step="any" />
				</div>
				<button
					onclick={locate}
					class="mt-1 w-fit text-xs text-muted-foreground underline underline-offset-2 transition-colors hover:text-foreground"
				>
					Use my current location
				</button>
			</div>

			{#if error}
				<p class="text-xs text-destructive">{error}</p>
			{/if}
		</div>

		<Dialog.Footer>
			<Button variant="secondary" onclick={() => (open = false)}>Cancel</Button>
			<Button onclick={handleSubmit} disabled={loading}>
				{loading ? 'Creating…' : 'Create'}
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
