<script lang="ts">
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { Button } from '$lib/components/ui/button';
	import type { AssetUrls, Location, StatusResponse } from '$lib/types';
	import { onMount } from 'svelte';
	import { safeFly } from '@/transitions.js';
	import type { Post } from '$lib/types';

	let ready = $state(false);

	onMount(() => (ready = true));

	type Props = {
		data: { location: Location; posts: Post[]; status: StatusResponse; assetUrls: AssetUrls };
	};

	let { data }: Props = $props();
	const loc: Location = $derived(data.location);
	const posts: Post[] = $derived(data.posts);
	const status: StatusResponse = $derived(data.status);
	const assetUrls: AssetUrls = $derived(data.assetUrls);
	let currentSlide = $state(0);
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
						{#if assetUrls?.paths?.length > 0}
							<!-- Image slide -->
							<div class="h-full w-full">
								<img
									src={`/api/media${assetUrls.paths[currentSlide]}`}
									alt="Location photo {currentSlide + 1}"
									class="h-full w-full object-cover"
								/>
							</div>

							<!-- Prev button -->
							<button
								onclick={() =>
									(currentSlide =
										(currentSlide - 1 + assetUrls.paths.length) % assetUrls.paths.length)}
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
								onclick={() => (currentSlide = (currentSlide + 1) % assetUrls.paths.length)}
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
								{#each assetUrls.paths as _, i}
									<button
										onclick={() => (currentSlide = i)}
										class="h-1.5 w-1.5 rounded-full transition-colors {i === currentSlide
											? 'bg-primary-foreground'
											: 'bg-primary-foreground/30'}"
									/>
								{/each}
							</div>
						{:else}
							<!-- Empty state -->
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
						{/if}
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

					<div class="shrink-0 px-1">
						<p class="mb-1 text-xs font-bold text-primary-foreground/90 sm:text-sm md:text-xl">
							Busyness:
							{#if status?.averageBusyness != null}
								<span class="font-normal text-primary-foreground/75"
									>{status.averageBusyness}/5</span
								>
							{:else}
								<span class="font-normal text-primary-foreground/50 italic">No recent reports</span>
							{/if}
						</p>
						{#if status?.averageBusyness != null}
							<div
								class="h-2 w-full overflow-hidden rounded-full bg-primary-foreground/20 sm:h-2.5 md:h-3"
							>
								<div
									class="h-full rounded-full bg-primary-foreground/80 transition-all duration-500"
									style="width: {(status.averageBusyness / 5) * 100}%"
								></div>
							</div>
						{/if}
					</div>
				</div>

				<!-- Right Column: Reviews -->
				<div
					class="order-2 flex min-h-0 flex-col gap-2 rounded-xl bg-card p-2 sm:gap-3 sm:p-3 md:col-span-7 md:gap-4 md:rounded-[2rem] md:p-6"
				>
					<div class="flex min-h-0 flex-1 flex-col gap-2 overflow-hidden sm:gap-3 md:gap-4">
						{#if posts === null || posts.length === 0}
							<div
								class="flex flex-1 flex-col items-center justify-center gap-2 text-muted-foreground"
							>
								<svg
									class="h-8 w-8 opacity-40 sm:h-10 sm:w-10 md:h-14 md:w-14"
									xmlns="http://www.w3.org/2000/svg"
									fill="none"
									viewBox="0 0 24 24"
									stroke="currentColor"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="1.5"
										d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-3 3-3-3z"
									/>
								</svg>
								<p class="text-xs font-medium sm:text-sm md:text-base">No reviews yet</p>
								<p class="text-xs opacity-60">Be the first to leave one!</p>
							</div>
						{:else}
							{#each posts as post (post.id)}
								<div
									class="flex flex-1 items-start gap-2 rounded-xl bg-secondary p-2 sm:gap-3 sm:p-3 md:gap-4 md:p-4"
								>
									<div
										class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-border text-xs font-bold text-muted-foreground sm:h-10 sm:w-10 md:h-16 md:w-16 md:text-base"
									>
										{post.rating}
									</div>

									<div class="flex min-w-0 flex-1 flex-col gap-1 sm:gap-1.5">
										<!-- Rating stars + date -->
										<div class="flex items-center gap-2">
											<div class="flex gap-0.5">
												{#each Array(5) as _, i}
													<svg
														class="h-3 w-3 sm:h-3.5 sm:w-3.5 md:h-4 md:w-4 {i < post.rating
															? 'text-yellow-400'
															: 'text-border'}"
														viewBox="0 0 20 20"
														fill="currentColor"
													>
														<path
															d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
														/>
													</svg>
												{/each}
											</div>
											<span class="text-[10px] text-muted-foreground sm:text-xs">
												{new Date(post.createdAt).toLocaleDateString()}
											</span>
										</div>

										<!-- Description -->
										<p
											class="line-clamp-3 text-xs text-secondary-foreground sm:text-sm md:text-base"
										>
											{post.description}
										</p>
									</div>
								</div>
							{/each}
						{/if}
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
