<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index';
	import { Input } from '$lib/components/ui/input/index';
	import { Label } from '$lib/components/ui/label/index';
	import { Slider } from '$lib/components/ui/slider/index';
	import { Button } from '$lib/components/ui/button/index';

	import { navigationMenuTriggerStyle } from '$lib/components/ui/navigation-menu/navigation-menu-trigger.svelte';
	import { cn } from '$lib/utils';

	let location = $state('');
	let busyness = $state([0]);

	function handleCheckIn() {
		console.log('Checking in at:', location, 'with busyness:', busyness[0]);
		// Add API call logic here if needed
	}
</script>

<Dialog.Root>
	<Dialog.Trigger>
		{#snippet child({ props })}
			<button
				{...props}
				class="inline-flex items-center justify-center gap-1 rounded-md px-4 py-2 text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground focus:outline-none"
			>
				Check In
			</button>
		{/snippet}
	</Dialog.Trigger>
	<Dialog.Content class="bg-[#2D4336] border-none text-white max-w-md rounded-3xl p-8">
		<Dialog.Header>
			<Dialog.Title class="text-3xl font-bold text-center underline decoration-2 underline-offset-8 mb-8">
				Check In
			</Dialog.Title>
		</Dialog.Header>

		<div class="space-y-8">
			<div class="space-y-4">
				<Label for="location" class="text-xl font-bold">Location:</Label>
				<Input
					id="location"
					placeholder="Enter Location"
					bind:value={location}
					class="bg-[#D9D9D9] text-[#2D4336] placeholder:text-[#2D4336]/60 border-none rounded-full h-12 text-lg px-6"
				/>
			</div>

			<div class="space-y-4">
				<Label class="text-xl font-bold">Busyness:</Label>
				<div class="px-2">
					<Slider
						bind:value={busyness}
						max={5}
						step={0.5}
						class="w-full [&_[data-slot=slider-thumb]]:bg-[#C4D7B2] [&_[data-slot=slider-thumb]]:size-6 [&_[data-slot=slider-track]]:bg-[#D9D9D9] [&_[data-slot=slider-range]]:bg-[#C4D7B2]"
					/>
					<div class="flex justify-between mt-4 px-1 text-lg font-bold items-center">
						{#each [0, 1, 2, 3, 4, 5] as num}
							<span>{num}</span>
							{#if num !== 5}
								<span class="text-2xl">•</span>
							{/if}
						{/each}
					</div>
				</div>
			</div>

			<div class="flex justify-center pt-4">
				<Button
					onclick={handleCheckIn}
					class="bg-[#D9D9D9] hover:bg-[#D9D9D9]/90 text-[#2D4336] font-bold rounded-2xl h-12 px-12 text-xl"
				>
					Check In
				</Button>
			</div>
		</div>
	</Dialog.Content>
</Dialog.Root>
