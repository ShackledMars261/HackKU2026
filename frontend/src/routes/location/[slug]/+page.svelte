<script lang="ts">
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { Button } from '$lib/components/ui/button';
	import * as Carousel from '$lib/components/ui/carousel';
	import type { Location } from '$lib/types';

	let { data } = $props();
	const loc: Location = $derived(data.location);

	function scrollable(node: HTMLElement) {
		document.body.style.overflow = 'auto';
		document.body.style.position = 'static';
		return {
			destroy() {
				document.body.style.overflow = 'hidden';
				document.body.style.position = 'fixed';
			}
		};
	}
</script>

{#if loc}
	<div
		use:scrollable
		class="flex min-h-screen w-full justify-center bg-background px-2 pt-8 pb-8 sm:px-3 sm:pt-12 md:px-4 md:pt-16"
	>
		<div
			class="grid h-fit w-full grid-cols-1 gap-3 rounded-2xl bg-primary p-3 shadow-lg sm:gap-4 sm:p-4 md:grid-cols-12 md:gap-8 md:rounded-[2.5rem] md:p-8"
			style="max-width: min(180vh, 1800px);"
		>
			<!-- Left Column: Location Details -->
			<div class="order-1 flex flex-col gap-3 sm:gap-4 md:col-span-5 md:gap-6">
				<Carousel.Root class="w-full">
					<Carousel.Content>
						{#each Array(4) as _, i}
							<Carousel.Item>
								<div
									class="flex w-full items-center justify-center rounded-xl bg-secondary"
									style="height: clamp(140px, 35vw, 44vh);"
								>
									<svg
										class="h-8 w-8 text-muted-foreground/40 sm:h-12 sm:w-12 md:h-24 md:w-24"
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
							</Carousel.Item>
						{/each}
					</Carousel.Content>
					<Carousel.Previous
						class="left-2 border-primary-foreground/20 bg-primary-foreground/10 text-primary-foreground hover:bg-primary-foreground/20 hover:text-primary-foreground"
					/>
					<Carousel.Next
						class="right-2 border-primary-foreground/20 bg-primary-foreground/10 text-primary-foreground hover:bg-primary-foreground/20 hover:text-primary-foreground"
					/>
				</Carousel.Root>

				<!-- Info -->
				<div class="flex flex-col gap-1.5 px-1 sm:gap-2 md:gap-4">
					<h1 class="text-xl font-extrabold text-primary-foreground sm:text-2xl md:text-5xl">
						{loc.name}
					</h1>
					<p class="text-xs font-bold text-primary-foreground/90 sm:text-sm md:text-2xl">
						Address:
						<span class="font-normal text-primary-foreground/75">
							{loc.location.coordinates[1]}, {loc.location.coordinates[0]}
						</span>
					</p>
					<p class="text-xs font-bold text-primary-foreground/90 sm:text-sm md:text-2xl">
						Avg. Rating:
						<span class="font-normal text-primary-foreground/75">
							{loc.overallRating}/5
						</span>
					</p>
					<p class="text-xs font-bold text-primary-foreground/90 sm:text-sm md:text-2xl">
						Status:
						<span class="font-normal text-primary-foreground/75"></span>
					</p>
				</div>

				<div>
					<Skeleton class="h-5 w-full rounded-full bg-primary-foreground/20 sm:h-6 md:h-10 md:w-4/5" />
				</div>
			</div>

			<!-- Right Column: Reviews -->
			<div
				class="order-2 flex flex-col gap-3 rounded-xl bg-card p-3 sm:gap-4 sm:p-4 md:col-span-7 md:gap-6 md:rounded-[2rem] md:p-8"
			>
				<div class="flex flex-col gap-2 sm:gap-3 md:gap-6">
					{#each Array(4) as _, i}
						<div class="flex items-center gap-2 rounded-xl bg-secondary p-2 sm:gap-4 sm:p-3 md:gap-6 md:p-6">
							<Skeleton class="h-8 w-8 shrink-0 rounded-full bg-border sm:h-12 sm:w-12 md:h-24 md:w-24" />
							<div class="flex w-full flex-col gap-1.5 sm:gap-2 md:gap-3">
								<Skeleton class="h-2.5 w-20 rounded bg-border sm:h-3 sm:w-24 md:h-5 md:w-36" />
								<Skeleton class="h-4 w-full rounded bg-border sm:h-5 md:h-10 md:w-[90%]" />
								<Skeleton class="h-2.5 w-3/4 rounded bg-border sm:h-3 md:h-5" />
							</div>
						</div>
					{/each}
				</div>

				<div class="mt-1 flex justify-end gap-2 md:mt-2 md:gap-4">
					<Button
						variant="secondary"
						class="h-8 rounded-full border border-border bg-secondary px-3 text-xs font-semibold text-secondary-foreground hover:bg-muted sm:h-10 sm:px-5 sm:text-sm md:h-16 md:px-10 md:text-xl"
					>
						View More
					</Button>
					<Button
						class="h-8 rounded-full px-3 text-xs font-semibold sm:h-10 sm:px-5 sm:text-sm md:h-16 md:px-10 md:text-xl"
					>
						Add Review
					</Button>
				</div>
			</div>
		</div>
	</div>
{/if}