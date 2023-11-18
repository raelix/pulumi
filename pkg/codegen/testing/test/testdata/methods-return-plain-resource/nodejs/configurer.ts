// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import * as pulumiTls from "@pulumi/tls";

export class Configurer extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'metaprovider:index:Configurer';

    /**
     * Returns true if the given object is an instance of Configurer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Configurer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Configurer.__pulumiType;
    }


    /**
     * Create a Configurer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ConfigurerArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["tlsProxy"] = args ? args.tlsProxy : undefined;
        } else {
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Configurer.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }

    meaningOfLife(): Promise<number> {
        return utilities.callAsync("metaprovider:index:Configurer/meaningOfLife", {
            "__self__": this,
        }, this, {property: "res"});    }

    objectMix(): Promise<Configurer.ObjectMixResult> {
        return utilities.callAsync("metaprovider:index:Configurer/objectMix", {
            "__self__": this,
        }, this, {});    }

    tlsProvider(): Promise<pulumiTls.Provider> {
        return utilities.callAsync("metaprovider:index:Configurer/tlsProvider", {
            "__self__": this,
        }, this, {property: "res"});    }
}

/**
 * The set of arguments for constructing a Configurer resource.
 */
export interface ConfigurerArgs {
    tlsProxy?: pulumi.Input<string>;
}

export namespace Configurer {
    /**
     * The results of the Configurer.objectMix method.
     */
    export interface ObjectMixResult {
        readonly meaningOfLife?: number;
        readonly provider?: pulumiTls.Provider;
    }

}
