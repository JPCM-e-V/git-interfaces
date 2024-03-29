package redisrepo

// import (
// 	// "container/list"
// 	"context"
// 	"errors"
// 	"fmt"
// 	"strings"

// 	"github.com/go-redis/redis/v8"
// )

// var redisclient *redis.Client

// const REPO_KEY string = "GIT_REPO"

// func Init() {
// 	redisclient = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
// }

// func LsRefs(repoName string) ([]string, error) {
// 	var repoExistsCmd = redisclient.SIsMember(context.Background(), fmt.Sprintf("%s:repos", REPO_KEY), repoName)
// 	if repoExistsCmd.Err() != nil && repoExistsCmd.Err() != redis.Nil {
// 		return nil, repoExistsCmd.Err()
// 	} else if repoExistsCmd.Err() == redis.Nil || !repoExistsCmd.Val() {
// 		return nil, errors.New("repo does not exist")
// 	}

// 	var refnamesCmd = redisclient.Keys(context.Background(), fmt.Sprintf("%s:repo:%s:ref:*", REPO_KEY, repoName))
// 	if refnamesCmd.Err() != nil {
// 		return nil, refnamesCmd.Err()
// 	}
// 	var refKeys = refnamesCmd.Val()
// 	var refObjs map[string]map[string]string = make(map[string]map[string]string)
// 	for _, refName := range refKeys {
// 		var refCmd = redisclient.HGetAll(context.Background(), refName)
// 		fmt.Printf("%s:repo:%s:ref:%s\n", REPO_KEY, repoName, refName)
// 		if refCmd.Err() != nil {
// 			return nil, refCmd.Err()
// 		}
// 		fmt.Print(refCmd.Val())
// 		refObjs[refName[strings.LastIndex(refName, ":")+1:]] = refCmd.Val()
// 	}
// 	for _, refObj := range refObjs {
// 		if obj_id, error := resolveObjId(refObjs, refObj); error != nil {
// 			return nil, error
// 		} else {
// 			refObj["obj-id"] = obj_id
// 		}
// 	}
// 	var refs = make([]string, len(refObjs))
// 	var i = 0
// 	for refName, refObj := range refObjs {
// 		var ref = refObj["obj-id"] + " " + refName
// 		if symrefTarget, ok := refObj["symref-target"]; ok {
// 			ref += " symref-target:" + symrefTarget
// 		}
// 		refs[i] = ref
// 		i++
// 	}
// 	return refs, nil
// }

// func resolveObjId(refObjs map[string]map[string]string, refObj map[string]string) (string, error) {
// 	if refType, ok := refObj["type"]; ok {
// 		switch refType {
// 		case "commit":
// 			if objId, ok := refObj["obj-id"]; ok {
// 				return objId, nil
// 			} else {
// 				return "", errors.New("missing key 'obj-id' in reference of type 'commit'")
// 			}
// 		case "symref":
// 			if objId, ok := refObj["obj-id"]; ok {
// 				return objId, nil
// 			}
// 			if symrefTarget, ok := refObj["symref-target"]; ok {
// 				if symrefTargetObj, ok := refObjs[symrefTarget]; ok {
// 					return resolveObjId(refObjs, symrefTargetObj)
// 				} else {
// 					return "", errors.New(fmt.Sprintf("invalid 'symref-target': '%s'", symrefTarget))
// 				}
// 			} else {
// 				return "", errors.New("missing key 'symref-target' in reference of type 'symref'")
// 			}
// 		default:
// 			return "", errors.New(fmt.Sprintf("unsupported type '%s'", refType))
// 		}
// 	} else {
// 		return "", errors.New("missing key 'type'")
// 	}
// }
