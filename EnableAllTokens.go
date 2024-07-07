package main

import (
    "fmt"
    "golang.org/x/sys/windows"
)

var tokens = []string{
    "SeAssignPrimaryTokenPrivilege",
    "SeAuditPrivilege",
    "SeBackupPrivilege",
    "SeChangeNotifyPrivilege",
    "SeCreateGlobalPrivilege",
    "SeCreatePagefilePrivilege",
    "SeCreatePermanentPrivilege",
    "SeCreateSymbolicLinkPrivilege",
    "SeCreateTokenPrivilege",
    "SeDebugPrivilege",
    "SeDelegateSessionUserImpersonatePrivilege",
    "SeEnableDelegationPrivilege",
    "SeImpersonatePrivilege",
    "SeIncreaseQuotaPrivilege",
    "SeIncreaseBasePriorityPrivilege",
    "SeIncreaseWorkingSetPrivilege",
    "SeLoadDriverPrivilege",
    "SeLockMemoryPrivilege",
    "SeMachineAccountPrivilege",
    "SeManageVolumePrivilege",
    "SeProfileSingleProcessPrivilege",
    "SeRelabelPrivilege",
    "SeRemoteShutdownPrivilege",
    "SeRestorePrivilege",
    "SeSecurityPrivilege",
    "SeShutdownPrivilege",
    "SeSyncAgentPrivilege",
    "SeSystemtimePrivilege",
    "SeSystemEnvironmentPrivilege",
    "SeSystemProfilePrivilege",
    "SeTakeOwnershipPrivilege",
    "SeTcbPrivilege",
    "SeTimeZonePrivilege",
    "SeTrustedCredManAccessPrivilege",
    "SeUndockPrivilege",
}

func main() {
    hProcess := windows.CurrentProcess()
    var hToken windows.Token

    err := windows.OpenProcessToken(hProcess, windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &hToken)
    if err != nil {
        fmt.Println("Error opening process token:", err)
        return
    }
    defer hToken.Close()

    for _, token := range tokens {
        var luid windows.LUID
        err := windows.LookupPrivilegeValue(nil, windows.StringToUTF16Ptr(token), &luid)
        if err != nil {
            fmt.Printf("Error looking up privilege value for %s: %v\n", token, err)
            continue
        }

        tp := windows.Tokenprivileges{
            PrivilegeCount: 1,
            Privileges: [1]windows.LUIDAndAttributes{
                {Luid: luid, Attributes: windows.SE_PRIVILEGE_ENABLED},
            },
        }

        err = windows.AdjustTokenPrivileges(hToken, false, &tp, 0, nil, nil)
        if err != nil {
            fmt.Printf("Error adjusting token privileges for %s: %v\n", token, err)
            continue
        }

        if windows.GetLastError() == windows.ERROR_NOT_ALL_ASSIGNED {
            fmt.Printf("The privilege %s was not assigned.\n", token)
        } else {
            fmt.Printf("The privilege %s was successfully adjusted.\n", token)
        }
    }
	fmt.Scanln()
}
