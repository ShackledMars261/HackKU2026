<script lang="ts">
	import * as Tabs from '$lib/components/ui/tabs/index';
	import * as Card from '$lib/components/ui/card/index';
	import { Button } from '$lib/components/ui/button/index';
	import { Input } from '$lib/components/ui/input/index';
	import { Label } from '$lib/components/ui/label/index';
	import { enhance } from '$app/forms';

	let username = $state('');
	let password = $state('');
</script>

<svelte:head>
    <title>SpotDrop</title>
</svelte:head>

<div class="flex h-[75%] w-full items-center justify-center bg-background text-sm sm:text-base md:text-lg lg:text-xl">
	<div class="-mb-4 flex w-full max-w-sm sm:max-w-md md:max-w-lg flex-col gap-6">
		<Tabs.Root value="login">
			<Tabs.List class="text-sm sm:text-base md:text-lg">
				<Tabs.Trigger value="login" class="text-sm sm:text-base md:text-lg px-3 sm:px-4 md:px-6">Login</Tabs.Trigger>
				<Tabs.Trigger value="signup" class="text-sm sm:text-base md:text-lg px-3 sm:px-4 md:px-6">Sign Up</Tabs.Trigger>
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
								<Label for="tabs-username" class="text-sm sm:text-base md:text-lg">Username</Label>
								<Input
									name="username"
									id="tabs-username"
									placeholder="Username"
									class="text-sm sm:text-base md:text-lg h-9 sm:h-10 md:h-11"
									bind:value={username}
								/>
							</div>
							<div class="grid gap-3">
								<Label for="tabs-password" class="text-sm sm:text-base md:text-lg">Password</Label>
								<Input
									name="password"
									id="tabs-password"
									placeholder="Password"
									class="text-sm sm:text-base md:text-lg h-9 sm:h-10 md:h-11"
									bind:value={password}
								/>
							</div>
						</Card.Content>
						<Card.Footer>
							<Button type="submit" class="text-sm sm:text-base md:text-lg px-4 sm:px-6 md:px-8 h-9 sm:h-10 md:h-11">Login</Button>
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
								Create a new account. If you already have an account, you can login by clicking the
								Login tab.
							</Card.Description>
						</Card.Header>
						<Card.Content class="grid gap-6">
							<div class="grid gap-3">
								<Label for="tabs-signup-username" class="text-sm sm:text-base md:text-lg">Username</Label>
								<Input
									name="username"
									id="tabs-signup-username"
									placeholder="Username"
									class="text-sm sm:text-base md:text-lg h-9 sm:h-10 md:h-11"
									value={username}
								/>
							</div>
							<div class="grid gap-3">
								<Label for="tabs-signup-password" class="text-sm sm:text-base md:text-lg">Password</Label>
								<Input
									name="password"
									id="tabs-signup-password"
									placeholder="Password"
									class="text-sm sm:text-base md:text-lg h-9 sm:h-10 md:h-11"
									value={password}
								/>
							</div>
						</Card.Content>
						<Card.Footer>
							<Button type="submit" class="text-sm sm:text-base md:text-lg px-4 sm:px-6 md:px-8 h-9 sm:h-10 md:h-11">Sign Up</Button>
						</Card.Footer>
					</Card.Root>
				</form>
			</Tabs.Content>
		</Tabs.Root>
	</div>
</div>