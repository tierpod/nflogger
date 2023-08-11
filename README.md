# nflogger

*nflogger* is a simple tool to dump packets logged via the Linux firewall NFLOG target.

Example rule and 

```
sudo nft add ip filter INPUT log prefix \"filter/INPUT: \" group 100
sudo nflogger --group 100
```

When you are done press CTRL-C to stopp nflogger and remove the rule again:

```
sudo nft -a list chain filter INPUT
```

This prints all rules of the chain INPUT including the handle id.

```
sudo nft delete rule filter INPUT handle <handle>
```

## License

    3-clause BSD

    © 2023 Christian Pointner <equinox@spreadspace.org>
