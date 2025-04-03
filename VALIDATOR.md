# Validator Reactivation

Description:

As a result of certain validator actions, its status changes to "Offline". To allow the validator to be re-added to the active validator list, it must change its status on the contract. However, no functionality other than validator creation and deactivation that changes the status was found.

Thus, the only solution is:

• Withdraw the deposited amount along with accumulated rewards (if they were not collected earlier),

• Subtract the inactivity penalty,

• Re-deposit the funds using a newly generated validator public address.

Problem:

Validator reactivation is essentially the process of creating a new validator. Since the contract does not have a function to delete invalid validators, there is a high theoretical risk that new validators cannot be created. This could potentially lead to network failure due to the inability to reactivate previously created validators.

## Validator Reactivation Process

When a validator goes offline or becomes inactive, follow these steps to reactivate it:

### Step 1: Withdraw Staked Tokens

From the validator address that needs reactivation, initiate a withdrawal request:

- Call the `undelegate` method on the SFC contract
- This will prepare your request to withdraw the native network tokens
- Initialize the SFC contract ABI variable found [here](abi/sfcc.js)
  ```bash
     # Attach to opera console
     (validator)$ ./opera attach opera.ipc
     #Parse to javascript terminal everything from abi.sfcc.js
     sfcc = web3.ftm.contract(abi).at("0xfc00face00000000000000000000000000000000")
  ```
- Unlock the wallet
  ```bash
    personal.unlockAccount("{VALIDATOR_WALLET_ADDRESS}", "{PASSWORD}", 60)
   ```
- Call the `undelegate` method
  ```bash
     # Get ValidatorID
     # Get ValidatorID from sfcc.getValidatorID("{VALIDATOR_WALLET_ADDRESS}") e.g 9
     # Amount e.g web3.toWei("200000.0", "ftm")
     # Call the undelegate method
     tx = sfcc.undelegate({ValidatorID}, {amount})
     # This returns a withdrawal request ID (wrID)
  ```

### Step 2: Wait for the Withdrawal Period
- The current withdrawal delay period is set to 1 hour (defined by `withdrawalPeriodTime`)
- This waiting period is a security measure to prevent immediate withdrawals

### Step 3: Complete the Withdrawal
- After the waiting period expires, call the `withdraw` method
- You'll need the withdrawal request ID (wrID) which can be obtained from the `undelegate` event
- This step completes the token withdrawal process
```bash
     # Attach to opera console
     (validator)$ ./opera attach opera.ipc
     #Parse to javascript terminal everything from abi.sfcc.js
     sfcc = web3.ftm.contract(abi).at("0xfc00face00000000000000000000000000000000")
  ```
- Unlock the wallet
  ```bash
    personal.unlockAccount("{VALIDATOR_WALLET_ADDRESS}", "{PASSWORD}", 60)
   ```
- Call the `withdraw` method
  ```bash
     # Get ValidatorID
     # Get ValidatorID from sfcc.getValidatorID("{VALIDATOR_WALLET_ADDRESS}") e.g 9
     # Amount e.g web3.toWei("200000.0", "ftm")
     # Call the withdraw method
     tx = sfcc.withdraw({ValidatorID}, {wrID})
  ```

### Step 4: Stop the Inactive Validator
- Properly shut down the validator node that was running
- This ensures a clean state before restarting with a new identity

### Step 5: Restart the Validator
- Generate a new validator key pair following the documentation
- Set up the validator node with the new keys
- Deposit the required stake (minus any penalties incurred)
- Complete the validator registration process

