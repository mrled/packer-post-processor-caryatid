# Dev Notes

Miscellaneous notes to self I made during this process

## Packer artifact flow

In designing this plugin, it was helpful for me to outline how artifacts go from one step to another in a Packer workflow

- Build using one or more relevant builders.
    - This should be something that makes sense for Vagrant to consume, such as Virtualbox.
    - Note that each builder will result in a single *artifact* that is composed of one or more *files*. For instance, the VMware builder outputs at least one `.vmdk` disk file as well as a `.vmx` file describing the virtual machine, and perhaps others as well.
- Optionally run one or more provisioners, which do not result in different artifacts, but instead allow packer to modify the VMs by logging into them and applying configuration changes
- Run the Vagrant post-processor
    - This is necessary because Caryatid is, after all, intended for use with Vagrant
    - This will take an artifact potentially consisting of several files, do some Vagrant-related processing, and output an artifact consisting of just one file with a filename ending in `.box`.
    - Note that each post-processor gets run once per artifact, so if you have defined multiple builders (say, a Virtualbox builder and a VMware builder), you'll end up with two artifacts, and the Vagrant provisioner will run once per artifact.
    - Note that since Vagrant boxes are provider-specific, the Vagrant post-processor is hard-coded to understand various providers and those providers only. The supported providers are listed in the documentation: https://www.packer.io/docs/post-processors/vagrant.html
- Run the Caryatid post-processor
    - Just like Vagrant, if you have defined multiple builders, this post-processor will see one artifact per builder (and modified by the Vagrant post-processor)
    - Note that we used to rely on the default Vagrant output filename to determine provider information, but since 1.0 we have determined it from the box file itself, which is an archive containing a `metadata.json` file that has the provider.
    - Based on configuration information specified in the packerfile, this will find a Vagrant catalog, add an entry for the relevant version and provider, and copy the Vagrant box file to a standard location linked from that catalog.
    - Note that by default it (and all Packer post-provisioners) will delete the input artifacts. This may be undesirable, for instance if your Vagrant catalog is on the Internet and you want to keep a local cache of it, so make sure to specify this in your packerfile.

## See also

I've never written a Packer plugin or, indeed, a single line of Go, so I found these links useful, and I am tired of hunting through my history to find them again

- The [Vagrant docs on the .box format](https://www.vagrantup.com/docs/boxes/format.html) goes into detail on the "box metadata", that is, what I have been calling the "Vagrant catalog"
- [Custom post-processor development](https://www.packer.io/docs/extend/post-processor.html) - official Packer documentation
- [How the first-party Varant post-processor works](https://www.packer.io/docs/post-processors/vagrant.html) - we consume the output of this post-processor
- [packer-post-processor-ami-file](https://github.com/scopely/packer-post-processor-ami-file) - A very tiny post-processor, useful for reference, with tests
- [packer-post-processor-vhd](https://github.com/benwebber/packer-post-processor-vhd) - A more complex post-processor that does unit testing (not even integration testing) on its PostProcess() method, also useful for reference
- [Packer source for the PostProcessor interface](https://github.com/mitchellh/packer/blob/master/packer/post_processor.go)
- [Packer source for the first-party Atlas post-processor](https://github.com/mitchellh/packer/blob/master/post-processor/atlas/post-processor.go)
- [How to set up a self-hosted "vagrant cloud" with versioned, self-packaged vagrant boxes](https://github.com/hollodotme/Helpers/blob/master/Tutorials/vagrant/self-hosted-vagrant-boxes-with-versioning.md) shows you how to write Vagrant catalogs manually
- [Distributing Vagrant base boxes securely](http://chase-seibert.github.io/blog/2014/05/18/vagrant-authenticated-private-box-urls.html) also discusses creating Vagrant catalogs manually
- The [Amazon error handling documentation](http://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/handling-errors.html) can help make sense of weird AWS SDK errors

