<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index';
	import { Button } from '$lib/components/ui/button/index';
	import { Label } from '$lib/components/ui/label/index';
	import { Slider } from '$lib/components/ui/slider/index';

	type Props = { locationId: string };
	let { locationId }: Props = $props();

	let open = $state(false);
	let description = $state('');
	let rating = $state([3]);
	let loading = $state(false);
	let error = $state<string | null>(null);
	let success = $state(false);

	async function handleSubmit() {
		if (!description.trim()) {
			error = 'Please write a review.';
			return;
		}

		loading = true;
		error = null;

		try {
			const resp = await fetch('/api/review', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					location_id: locationId,
					rating: rating[0],
					description: description.trim(),
					photo_file_ids: []
				})
			});

			if (!resp.ok) throw new Error(await resp.text());

			success = true;
			setTimeout(() => {
				open = false;
				description = '';
				rating = [3];
				success = false;
			}, 1000);
		} catch (e: any) {
			error = e.message;
		} finally {
			loading = false;
		}
	}

	function onClose() {
		if (loading) return;
		open = false;
		description = '';
		rating = [3];
		error = null;
		success = false;
	}
</script>

<Dialog.Root bind:open>
	<Dialog.Trigger>
		{#snippet child({ props })}
			<Button
				{...props}
				class="h-8 rounded-full px-3 text-xs font-semibold sm:h-9 sm:px-4 sm:text-sm md:h-12 md:px-8 md:text-base"
			>
				Add Review
			</Button>
		{/snippet}
	</Dialog.Trigger>

	<Dialog.Content class="sm:max-w-md">
		<Dialog.Header>
			<Dialog.Title>Write a Review</Dialog.Title>
			<Dialog.Description>Share your experience at this location.</Dialog.Description>
		</Dialog.Header>

		<div class="flex flex-col gap-5 py-2">
			<!-- Rating -->
			<div class="flex flex-col gap-3">
				<div class="flex items-center justify-between">
					<Label>Rating</Label>
					<span class="text-sm font-bold text-foreground">{rating[0]} / 5</span>
				</div>
				<Slider bind:value={rating} min={1} max={5} step={1} class="w-full" />
				<div class="flex justify-between px-1 text-xs text-muted-foreground">
					{#each [1, 2, 3, 4, 5] as n}
						<span>{n}</span>
					{/each}
				</div>
			</div>

			<!-- Description -->
			<div class="flex flex-col gap-1.5">
				<Label for="review-description">Review</Label>
				<textarea
					id="review-description"
					bind:value={description}
					placeholder="Write your review here…"
					rows={4}
					class="w-full resize-none rounded-md border border-input bg-background px-3 py-2 text-sm text-foreground placeholder:text-muted-foreground focus:ring-1 focus:ring-ring focus:outline-none"
				></textarea>
			</div>

			{#if success}
				<p class="flex items-center gap-1.5 text-sm text-green-600">
					<svg class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
						<path
							fill-rule="evenodd"
							d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
							clip-rule="evenodd"
						/>
					</svg>
					Review submitted!
				</p>
			{/if}

			{#if error}
				<p class="text-sm text-destructive">{error}</p>
			{/if}
		</div>

		<Dialog.Footer>
			<Button variant="secondary" onclick={onClose} disabled={loading}>Cancel</Button>
			<Button onclick={handleSubmit} disabled={loading || !description.trim()}>
				{loading ? 'Submitting…' : 'Submit'}
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
