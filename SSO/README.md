# Unique User Management system

## Function

1. SSO

    the redirect is done by the traefik gateway. So any request sent to the real service is an authenticated request. The SSO will insert 2 http header:
    
    - `X-ID`: contains the user id

2. Permission control

    the permission is divided by 3 parts:
    
    - `Action`: `READ`, `WRITE`, `EXECUTE`
    - `Resource`: `SMS`, `EMAIL` and so on
    - `Scope`: `SELF`, `GROUP`, `ALL`
    
    For third-party application, before any instructions related to specific `Resource`, they must ask the User System to query whether the user has the permission to do it. The real world meaning about the permission is defined in `resource-actions.md`

3. Staff Turnover

    Always, we are shamed to remove some regretful newbies from our group. Also it's hard to transfer the power when the leadership changes. Therefore, let program help you to do it.

    We supports below actions: 

    - kick the underproof from the QQ group, Telegram group and the lark organization
        > Of course, you need to introduce our bot into the group first.
    - out-of-box leadership change. Free the permission of the older while authorize the newer for our system and the lark.