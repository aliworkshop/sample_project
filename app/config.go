package app

import "time"

const ServiceName = "bc-client"

type config struct {
	Http struct {
		Address                   string
		GracefullyShutdownTimeout time.Duration
	}
}

func (c *config) Initialize() {
	if c.Http.GracefullyShutdownTimeout == 0 {
		c.Http.GracefullyShutdownTimeout = time.Second * 10
	}
}

const (
	ScopePortfoliosReadAll           = "api.wallet.portfolios.read-all"
	ScopeWithdrawRequestsReadAll     = "api.wallet.withdraw-requests.read-all"
	ScopeWithdrawRequestVerifiesRead = "api.wallet.withdraw-requests.verifies.read"
	ScopeWithdrawRequestVerify       = "api.wallet.withdraw-requests.verify"
	ScopeBcTxReadAll                 = "api.wallet.blockchain-transactions.read-all"
	ScopeDashboardRead               = "api.wallet.dashboard.read"
	ScopePortfolioHistoriesReadAll   = "api.wallet.portfolio-histories.read-all"
	ScopePortfolioSummariesReadAll   = "api.wallet.portfolio-summaries.read-all"
	ScopeTemplatesReload             = "api.wallet.templates.reload"

	ScopeBcTxRead               = "api.wallet.blockchain-transactions.read"
	ScopePortfoliosRead         = "api.wallet.portfolios.read"
	ScopePortfolioHistoriesRead = "api.wallet.portfolio-histories.read"
	ScopePortfolioSummaryRead   = "api.wallet.portfolio-summary.read"
	ScopeDepositRequest         = "api.wallet.deposit.request"
	ScopeWithdrawRequest        = "api.wallet.withdraw.request"
	ScopeWithdrawRequestsRead   = "api.wallet.withdraw-requests.read"
	ScopeBlockchainLoadBlock    = "api.bc-client.network.load-block"
	ScopeTetherlandWithdraw     = "api.bc-client.convert.withdraw"
	ScopePaymentWithdraw        = "api.bc-client.payment.withdraw"
)
