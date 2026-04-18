<script lang="ts">
	import * as NavigationMenu from '$lib/components/ui/navigation-menu/index';
	import { navigationMenuTriggerStyle } from '$lib/components/ui/navigation-menu/navigation-menu-trigger.svelte';
	import { IsMobile } from '$lib/hooks/is-mobile.svelte';
	import { resolve } from '$app/paths';

	const isMobile = new IsMobile();
</script>

<NavigationMenu.Root viewport={isMobile.current}>
	<NavigationMenu.List class="flex-wrap">
		<NavigationMenu.Item>
			<NavigationMenu.Link>
				{#snippet child()}
					<a href={resolve('/')} class={navigationMenuTriggerStyle()}>Home</a>
				{/snippet}
			</NavigationMenu.Link>
		</NavigationMenu.Item>
	</NavigationMenu.List>
</NavigationMenu.Root>
