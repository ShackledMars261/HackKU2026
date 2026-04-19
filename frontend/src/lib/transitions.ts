import { fly } from 'svelte/transition';
import type { FlyParams } from 'svelte/transition';

function noopTransition() {
	return { duration: 0 };
}

function prefersReducedMotion(): boolean {
	if (typeof window === 'undefined') return false;
	return window.matchMedia('(prefers-reduced-motion: reduce)').matches;
}

export function safeFly(node: Element, params: FlyParams) {
	if (prefersReducedMotion()) return noopTransition();
	return fly(node, params);
}
