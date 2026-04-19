<script>
	import heroBg from '$lib/assets/hero-bg.svg';
</script>

<svelte:head>
    <title>SpotDrop</title>
</svelte:head>


<div class="w-full min-h-screen flex flex-col bg-background overflow-scroll">
	
	<!-- HERO -->
	<div
		 style="background-image: url({heroBg}); background-repeat: no-repeat;  background-position: center 70px"
		class="bg-cover bg-top w-full min-h-screen flex sm:justify-center sm:items-center">
		<div
			class="flex flex-col justify-center items-center gap-3
			w-full h-full
			bg-card
			px-6 py-10 text-sm font-medium text-muted-foreground
			tracking-wide text-center

			sm:max-w-2xl sm:mx-auto sm:h-auto
			sm:border sm:border-border sm:rounded-3xl
			sm:px-12 sm:py-12 sm:shadow-sm"
		>
			<h1 class="text-5xl md:text-6xl font-bold text-foreground leading-tight tracking-tight mb-6">
				Study where it's<br/>
				<span class="italic text-accent">actually quiet.</span>
			</h1>

			<p class="text-lg md:text-xl text-muted-foreground font-light max-w-xl leading-relaxed mb-12">
				Discover cafés, libraries, and hidden corners near you sorted by how busy they are right now, not just on a good day.
			</p>

			<div class="flex flex-col sm:flex-row gap-3 items-center">
				<a href="/dashboard" class="bg-primary hover:bg-primary/90 text-primary-foreground text-sm font-medium px-7 py-3.5 rounded-full transition-colors shadow-md">
					Find a quiet spot
				</a>
				<a href="#about" class="text-muted-foreground hover:text-foreground text-sm font-medium px-5 py-3.5 transition-colors">
					How it works ↓
				</a>
			</div>
		</div>
	</div>

	<!-- NEXT SECTION -->
	<div id="about" class="w-full min-h-screen flex justify-center items-center bg-foreground">
		<div
			class="flex flex-col justify-center items-center gap-3
			w-full px-6 py-10 text-sm font-medium text-muted-foreground text-center

			sm:max-w-2xl sm:mx-auto
			sm:border sm:border-border sm:rounded-3xl
			sm:px-12 sm:py-12 sm:shadow-sm"
		>
			Content here
		</div>
	</div>

</div>