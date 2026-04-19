<script>
	import heroBg from '$lib/assets/hero-bg.svg';
</script>

<svelte:head>
	<title>SpotDrop</title>
</svelte:head>

<div class="flex min-h-screen w-full flex-col overflow-scroll bg-background">
	<!-- HERO -->
	<div
		style="background-image: url({heroBg}); background-repeat: no-repeat;  background-position: center 70px"
		class="flex min-h-screen w-full bg-cover bg-top sm:items-center sm:justify-center"
	>
		<div
			class="flex h-full w-full flex-col items-center
			justify-center gap-3
			bg-card
			px-6 py-10 text-center text-sm font-medium
			tracking-wide text-muted-foreground

			sm:mx-auto sm:h-auto sm:max-w-2xl
			sm:rounded-3xl sm:border sm:border-border
			sm:px-12 sm:py-12 sm:shadow-sm"
		>
			<h1 class="mb-6 text-5xl leading-tight font-bold tracking-tight text-foreground md:text-6xl">
				Study where it's<br />
				<span class="text-accent italic">actually quiet.</span>
			</h1>

			<p class="mb-12 max-w-xl text-lg leading-relaxed font-light text-muted-foreground md:text-xl">
				Discover cafés, libraries, and hidden corners near you sorted by how busy they are right
				now, not just on a good day.
			</p>

			<div class="flex flex-col items-center gap-3 sm:flex-row">
				<a
					href="/dashboard"
					class="rounded-full bg-primary px-7 py-3.5 text-sm font-medium text-primary-foreground shadow-md transition-colors hover:bg-primary/90"
					data-sveltekit-reload
				>
					Find a quiet spot
				</a>
				<a
					href="#about"
					class="px-5 py-3.5 text-sm font-medium text-muted-foreground transition-colors hover:text-foreground"
				>
					How it works ↓
				</a>
			</div>
		</div>
	</div>

	<!-- NEXT SECTION -->
	<div id="about" class="flex min-h-screen w-full items-center justify-center bg-foreground">
		<div
			class="flex w-full flex-col items-center justify-center
			gap-3 px-6 py-10 text-center text-sm font-medium text-muted-foreground

			sm:mx-auto sm:max-w-2xl
			sm:rounded-3xl sm:border sm:border-border
			sm:px-12 sm:py-12 sm:shadow-sm"
		>
			Content here
		</div>
	</div>
</div>
