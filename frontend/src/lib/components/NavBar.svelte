<script lang="ts">
	import * as NavigationMenu from '$lib/components/ui/navigation-menu/index';
	import { navigationMenuTriggerStyle } from '$lib/components/ui/navigation-menu/navigation-menu-trigger.svelte';
	import { IsMobile } from '$lib/hooks/is-mobile.svelte';
	import { resolve } from '$app/paths';
	import logo from '$lib/assets/logo.svg';

	const isMobile = new IsMobile();

	type Props = {
		isSignedIn: boolean;
	};

	let { isSignedIn }: Props = $props();
</script>

<div class="flex w-full items-center justify-between text-foreground">
	<NavigationMenu.Root viewport={isMobile.current} class="w-max max-w-max">
		<NavigationMenu.List>
			<NavigationMenu.Item>
				<NavigationMenu.Link>
					{#snippet child()}
						<a href={resolve('/')} class={navigationMenuTriggerStyle()} aria-label="Home">
							<img src={logo} alt="Home" class="w-30" />
						</a>
					{/snippet}
				</NavigationMenu.Link>
			</NavigationMenu.Item>
		</NavigationMenu.List>
	</NavigationMenu.Root>

	<NavigationMenu.Root viewport={isMobile.current} class="w-max max-w-max">
		<NavigationMenu.List>
			<NavigationMenu.Item>
				<NavigationMenu.Link>
					{#snippet child()}
						<a
							href={resolve('/dashboard')}
							class={navigationMenuTriggerStyle()}
							aria-label="Dashboard"
							data-sveltekit-reload
						>
							Dashboard
						</a>
					{/snippet}
				</NavigationMenu.Link>
			</NavigationMenu.Item>
			{#if !isSignedIn}
				<NavigationMenu.Item>
					<NavigationMenu.Link>
						{#snippet child()}
							<a
								href={resolve('/auth')}
								class={navigationMenuTriggerStyle()}
								aria-label="Login"
								data-sveltekit-reload
							>
								Login
							</a>
						{/snippet}
					</NavigationMenu.Link>
				</NavigationMenu.Item>
			{/if}
		</NavigationMenu.List>
	</NavigationMenu.Root>
</div>
