<svelte:head>
	<title>SpotDrop</title>
</svelte:head>
<div class="flex h-screen w-full items-center justify-center bg-background">
	<div
		class="mx-auto flex max-w-2xl flex-col items-center justify-center gap-2 rounded-3xl border border-border bg-card/70 px-16 py-16 text-center text-xs font-medium tracking-wide text-muted-foreground shadow-sm"
	>
		<div>
			<h1 class="mb-6 text-5xl leading-tight font-bold tracking-tight text-foreground md:text-6xl">
				Study where it's<br />
				<span class="text-accent italic">actually quiet.</span>
			</h1>
		</div>
		<p class="mb-12 max-w-xl text-lg leading-relaxed font-light text-muted-foreground md:text-xl">
			Discover cafés, libraries, and hidden corners near you sorted by how busy they are right now,
			not just on a good day.
		</p>
		<div class="flex flex-col items-center gap-3 sm:flex-row">
			<a
				href="#"
				class="rounded-full bg-primary px-7 py-3.5 text-sm font-medium text-primary-foreground shadow-md transition-colors hover:bg-primary/90"
			>
				Find a quiet spot
			</a>
			<a
				href="#"
				class="px-5 py-3.5 text-sm font-medium text-muted-foreground transition-colors hover:text-foreground"
			>
				How it works ↓
			</a>
		</div>
	</div>
</div>
