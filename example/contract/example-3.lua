-- Mints a new token
function Mint(minter, amount, symbol)
    set_value(minter, amount)
    set_value("minter", minter)
    set_value("symbol", symbol)
end

-- Transfer a token
function Transfer (sender, recipient, amount)
    -- Get sender balance
    senderBal = get_value(sender)
    -- Check sender exists
    if senderBal == "" then
        return "sender does not exist"
    -- Check sender has enough funds
    elseif senderBal < get_value(recipient) then
        return "sender does not have enough funds"
    end
    -- If recipient doesnt exists, set balance to 0
    if get_value(recipient) == "" then
        recipientBal = 0
    else
        -- If recipient exists set balance to current
        recipientBal = get_value(recipient)
    end
    -- Save new balances
    set_value(sender, senderBal - amount)
    set_value(recipient, recipientBal + amount)
    return "transfer complete"
 end

-- Returns the token owner
function getMinter ()
   return get_value("minter")
end

-- Returns the token symbol
function getSymbol ()
    return get_value("symbol")
end

-- Returns the balance of an account
function getAccount (account)
    return get_value(account)
 end

