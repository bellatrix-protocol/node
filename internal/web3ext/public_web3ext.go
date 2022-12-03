package web3ext

const PublicJs = `
web3._extend({
	property: 'public',
	methods: [
		new web3._extend.Method({
			name: 'getHeaderByNumber',
			call: 'public_getHeaderByNumber',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter]
		}),
		new web3._extend.Method({
			name: 'getHeaderByHash',
			call: 'eth_getHeaderByHash',
			params: 1
		}),
	]
});
`
