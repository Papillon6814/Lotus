<template>
	<div>
		<div class="container">
			<!-- this line is used by starport scaffolding # 4 -->
			<!-- <SpType modulePath="papillon6814.voter.voter" moduleType="Vote"  />
			<SpType modulePath="papillon6814.voter.voter" moduleType="Poll"  /> -->
			<PollForm />
			<PollList />
		</div>
	</div>
</template>

<script>
import PollForm from "../components/PollForm"
import PollList from "../components/PollList"

export default {
	name: 'Types',
	components: { PollForm, PollList }
}
</script>
