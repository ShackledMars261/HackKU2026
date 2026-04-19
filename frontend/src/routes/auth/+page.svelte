<script lang="ts">
	import * as Tabs from '$lib/components/ui/tabs/index';
	import * as Card from '$lib/components/ui/card/index';
	import { Button } from '$lib/components/ui/button/index';
	import { Input } from '$lib/components/ui/input/index';
	import { Label } from '$lib/components/ui/label/index';
	import { enhance } from '$app/forms';
	import { safeFly } from '@/transitions';
	import { onMount } from 'svelte';

	let ready = $state(false);

	let username = $state('');
	let password = $state('');

	onMount(() => (ready = true));
</script>

<svelte:head>
	<title>SpotDrop</title>
</svelte:head>

<div
	class="flex h-[75%] w-full items-center justify-center bg-background text-sm sm:text-base md:text-lg lg:text-xl"
>
	{#if ready}
		<div
			class="-mb-4 flex w-full max-w-sm flex-col gap-6 sm:max-w-md md:max-w-lg"
			in:safeFly={{ y: 300, duration: 1000 }}
		>
			<Tabs.Root value="login">
				<Tabs.List class="text-sm sm:text-base md:text-lg">
					<Tabs.Trigger value="login" class="px-3 text-sm sm:px-4 sm:text-base md:px-6 md:text-lg"
						>Login</Tabs.Trigger
					>
					<Tabs.Trigger value="signup" class="px-3 text-sm sm:px-4 sm:text-base md:px-6 md:text-lg"
						>Sign Up</Tabs.Trigger
					>
				</Tabs.List>
				<Tabs.Content value="login">
					<form method="POST" action="?/login" use:enhance>
						<Card.Root>
							<Card.Header>
								<Card.Title class="text-lg sm:text-xl md:text-2xl lg:text-3xl">Login</Card.Title>
								<Card.Description class="text-xs sm:text-sm md:text-base">
									Login to your account. If you don&apos;t have an account, you can create one by
									clicking the Sign Up tab.
								</Card.Description>
							</Card.Header>
							<Card.Content class="grid gap-6">
								<div class="grid gap-3">
									<Label for="tabs-username" class="text-sm sm:text-base md:text-lg">Username</Label
									>
									<Input
										name="username"
										id="tabs-username"
										placeholder="Username"
										class="h-9 text-sm sm:h-10 sm:text-base md:h-11 md:text-lg"
										bind:value={username}
									/>
								</div>
								<div class="grid gap-3">
									<Label for="tabs-password" class="text-sm sm:text-base md:text-lg">Password</Label
									>
									<Input
										name="password"
										id="tabs-password"
										placeholder="Password"
										class="h-9 text-sm sm:h-10 sm:text-base md:h-11 md:text-lg"
										bind:value={password}
									/>
								</div>
							</Card.Content>
							<Card.Footer>
								<Button
									type="submit"
									class="h-9 px-4 text-sm sm:h-10 sm:px-6 sm:text-base md:h-11 md:px-8 md:text-lg"
									>Login</Button
								>
							</Card.Footer>
						</Card.Root>
					</form>
				</Tabs.Content>
				<Tabs.Content value="signup">
					<form method="POST" action="?/register" use:enhance>
						<Card.Root>
							<Card.Header>
								<Card.Title class="text-lg sm:text-xl md:text-2xl lg:text-3xl">Sign Up</Card.Title>
								<Card.Description class="text-xs sm:text-sm md:text-base">
									Create a new account. If you already have an account, you can login by clicking
									the Login tab.
								</Card.Description>
							</Card.Header>
							<Card.Content class="grid gap-6">
								<div class="grid gap-3">
									<Label for="signup-username" class="text-sm sm:text-base md:text-lg"
										>Username</Label
									>
									<Input
										name="username"
										id="signup-username"
										placeholder="Username"
										class="h-9 text-sm sm:h-10 sm:text-base md:h-11 md:text-lg"
										value={username}
									/>
								</div>
								<div class="grid gap-3">
									<Label for="signup-password" class="text-sm sm:text-base md:text-lg"
										>Password</Label
									>
									<Input
										name="password"
										id="signup-password"
										placeholder="Password"
										class="h-9 text-sm sm:h-10 sm:text-base md:h-11 md:text-lg"
										value={password}
									/>
								</div>
							</Card.Content>
							<Card.Footer>
								<Button
									type="submit"
									class="h-9 px-4 text-sm sm:h-10 sm:px-6 sm:text-base md:h-11 md:px-8 md:text-lg"
									>Sign Up</Button
								>
							</Card.Footer>
						</Card.Root>
					</form>
				</Tabs.Content>
			</Tabs.Root>
		</div>
	{/if}
</div>
