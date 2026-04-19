<script lang="ts">
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { Button } from '$lib/components/ui/button';
	import type { Location } from '$lib/types';
	import { onMount } from 'svelte';
	import { safeFly } from '@/transitions.js';

	let ready = $state(false);
	let { data } = $props();
	const loc: Location = $derived(data.location);
	let currentSlide = $state(0);
	const totalSlides = 4;

	onMount(() => (ready = true));
</script>

{#if loc}
	<div
		class="flex h-[calc(100vh-var(--navbar-height,4rem))] w-screen items-center justify-center overflow-hidden bg-background p-2 sm:p-3 md:p-4"
	>
		{#if ready}
			<div
				class="grid h-full w-full grid-cols-1 gap-2 rounded-2xl bg-primary p-2 shadow-lg sm:gap-3 sm:p-3 md:grid-cols-12 md:gap-6 md:rounded-[2.5rem] md:p-6"
				style="max-width: min(180vh, 1800px);"
				in:safeFly={{ x: -200, duration: 600 }}
			>
				<!-- Left Column -->
				<div class="order-1 flex min-h-0 flex-col gap-2 sm:gap-3 md:col-span-5 md:gap-4">
					<!-- Manual Carousel -->
					<div class="relative min-h-0 flex-1 overflow-hidden rounded-xl bg-secondary">
						<!-- Slide -->
						<div class="flex h-full w-full items-center justify-center">
							<svg
								class="h-8 w-8 text-muted-foreground/40 sm:h-12 sm:w-12 md:h-16 md:w-16"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="1"
									d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
								/>
							</svg>
						</div>

						<!-- Prev button -->
						<button
							onclick={() => (currentSlide = (currentSlide - 1 + totalSlides) % totalSlides)}
							class="absolute top-1/2 left-2 flex h-7 w-7 -translate-y-1/2 items-center justify-center rounded-full border border-primary-foreground/20 bg-primary-foreground/10 text-primary-foreground hover:bg-primary-foreground/20 sm:h-8 sm:w-8"
						>
							<svg
								class="h-3 w-3 sm:h-4 sm:w-4"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M15 19l-7-7 7-7"
								/>
							</svg>
						</button>

						<!-- Next button -->
						<button
							onclick={() => (currentSlide = (currentSlide + 1) % totalSlides)}
							class="absolute top-1/2 right-2 flex h-7 w-7 -translate-y-1/2 items-center justify-center rounded-full border border-primary-foreground/20 bg-primary-foreground/10 text-primary-foreground hover:bg-primary-foreground/20 sm:h-8 sm:w-8"
						>
							<svg
								class="h-3 w-3 sm:h-4 sm:w-4"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M9 5l7 7-7 7"
								/>
							</svg>
						</button>

						<!-- Dots -->
						<div class="absolute bottom-2 left-1/2 flex -translate-x-1/2 gap-1.5">
							{#each Array(totalSlides) as _, i}
								<button
									onclick={() => (currentSlide = i)}
									class="h-1.5 w-1.5 rounded-full transition-colors {i === currentSlide
										? 'bg-primary-foreground'
										: 'bg-primary-foreground/30'}"
								/>
							{/each}
						</div>
					</div>

					<!-- Info -->
					<div class="flex shrink-0 flex-col gap-1 px-1 sm:gap-1.5 md:gap-2">
						<h1 class="text-lg font-extrabold text-primary-foreground sm:text-2xl md:text-4xl">
							{loc.name}
						</h1>
						<p class="text-xs font-bold text-primary-foreground/90 sm:text-sm md:text-xl">
							Address:
							<span class="font-normal text-primary-foreground/75">
								{loc.location.coordinates[1]}, {loc.location.coordinates[0]}
							</span>
						</p>
						<p class="text-xs font-bold text-primary-foreground/90 sm:text-sm md:text-xl">
							Avg. Rating:
							<span class="font-normal text-primary-foreground/75">
								{loc.overallRating}/5
							</span>
						</p>
						<p class="text-xs font-bold text-primary-foreground/90 sm:text-sm md:text-xl">
							Status:
							<span class="font-normal text-primary-foreground/75"></span>
						</p>
					</div>

					<div class="shrink-0">
						<Skeleton
							class="h-4 w-full rounded-full bg-primary-foreground/20 sm:h-5 md:h-8 md:w-4/5"
						/>
					</div>
				</div>

				<!-- Right Column: Reviews -->
				<div
					class="order-2 flex min-h-0 flex-col gap-2 rounded-xl bg-card p-2 sm:gap-3 sm:p-3 md:col-span-7 md:gap-4 md:rounded-[2rem] md:p-6"
				>
					<div class="flex min-h-0 flex-1 flex-col gap-2 overflow-hidden sm:gap-3 md:gap-4">
						{#each Array(4) as _, i}
							<div
								class="flex flex-1 items-center gap-2 rounded-xl bg-secondary p-2 sm:gap-3 sm:p-3 md:gap-4 md:p-4"
							>
								<Skeleton
									class="h-8 w-8 shrink-0 rounded-full bg-border sm:h-10 sm:w-10 md:h-16 md:w-16"
								/>
								<div class="flex w-full flex-col gap-1.5 sm:gap-2">
									<Skeleton class="h-2.5 w-20 rounded bg-border sm:h-3 sm:w-24 md:h-4 md:w-36" />
									<Skeleton class="h-3.5 w-full rounded bg-border sm:h-4 md:h-6 md:w-[90%]" />
									<Skeleton class="h-2.5 w-3/4 rounded bg-border sm:h-3 md:h-4" />
								</div>
							</div>
						{/each}
					</div>

					<div class="flex shrink-0 justify-end gap-2 md:gap-3">
						<Button
							variant="secondary"
							class="h-8 rounded-full border border-border bg-secondary px-3 text-xs font-semibold text-secondary-foreground hover:bg-muted sm:h-9 sm:px-4 sm:text-sm md:h-12 md:px-8 md:text-base"
						>
							View More
						</Button>
						<Button
							class="h-8 rounded-full px-3 text-xs font-semibold sm:h-9 sm:px-4 sm:text-sm md:h-12 md:px-8 md:text-base"
						>
							Add Review
						</Button>
					</div>
				</div>
			</div>
		{/if}
	</div>
{/if}
