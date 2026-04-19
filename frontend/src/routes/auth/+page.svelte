<script lang="ts">
	import { Button } from '$lib/components/ui/button/index';
	import { Input } from '$lib/components/ui/input/index';
	import { Label } from '$lib/components/ui/label/index';
	import { enhance } from '$app/forms';
	import { safeFly } from '@/transitions';
	import { onMount } from 'svelte';

	let ready = $state(false);
	let tab: 'login' | 'signup' = $state('login');

	onMount(() => (ready = true));
</script>

<svelte:head>
	<title>SpotDrop</title>
</svelte:head>

<div class="flex h-[75%] w-full items-center justify-center bg-background">
	{#if ready}
		<div class="w-full max-w-sm sm:max-w-md" in:safeFly={{ y: 300, duration: 1000 }}>
			<!-- Dialog-style card -->
			<div class="rounded-2xl border border-border bg-card p-6 shadow-lg sm:p-8">
				<!-- Header -->
				<div class="mb-6">
					<h1 class="text-lg font-semibold text-foreground sm:text-xl">
						{tab === 'login' ? 'Welcome back' : 'Create an account'}
					</h1>
					<p class="mt-1 text-sm text-muted-foreground">
						{tab === 'login'
							? 'Enter your credentials to sign in to your account.'
							: 'Enter your details below to create your account.'}
					</p>
				</div>

				<!-- Tab switcher -->
				<div class="mb-6 flex rounded-lg border border-border bg-secondary p-1">
					<button
						type="button"
						onclick={() => (tab = 'login')}
						class="flex-1 rounded-md py-1.5 text-sm font-medium transition-colors
							{tab === 'login'
							? 'bg-background text-foreground shadow-sm'
							: 'text-muted-foreground hover:text-foreground'}"
					>
						Login
					</button>
					<button
						type="button"
						onclick={() => (tab = 'signup')}
						class="flex-1 rounded-md py-1.5 text-sm font-medium transition-colors
							{tab === 'signup'
							? 'bg-background text-foreground shadow-sm'
							: 'text-muted-foreground hover:text-foreground'}"
					>
						Sign Up
					</button>
				</div>

				<!-- Login form -->
				{#if tab === 'login'}
					<form method="POST" action="?/login" use:enhance class="flex flex-col gap-4">
						<div class="flex flex-col gap-1.5">
							<Label for="login-username">Username</Label>
							<Input id="login-username" name="username" placeholder="Username" />
						</div>
						<div class="flex flex-col gap-1.5">
							<Label for="login-password">Password</Label>
							<Input id="login-password" name="password" type="password" placeholder="Password" />
						</div>
						<div class="mt-2 flex justify-end gap-2">
							<Button type="submit" class="w-full">Login</Button>
						</div>
					</form>

					<!-- Signup form -->
				{:else}
					<form method="POST" action="?/register" use:enhance class="flex flex-col gap-4">
						<div class="flex flex-col gap-1.5">
							<Label for="signup-username">Username</Label>
							<Input id="signup-username" name="username" placeholder="Username" />
						</div>
						<div class="flex flex-col gap-1.5">
							<Label for="signup-password">Password</Label>
							<Input id="signup-password" name="password" type="password" placeholder="Password" />
						</div>
						<div class="mt-2 flex justify-end gap-2">
							<Button type="submit" class="w-full">Sign Up</Button>
						</div>
					</form>
				{/if}
			</div>
		</div>
	{/if}
</div>
