vim.g.runnercommand = { "make" }
vim.g.runnerfiletypeexceptions = {
	"http",
}
vim.g.runnerpattern = { "go", "html" }
--
-- this is so extraneaous commands are not run unintentionally
vim.g.ignorepattern = { "main.go" }
